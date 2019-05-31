package auth

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ligenhw/goshare/user"
)

type AlipayClient struct {
	Url           string
	AppID         string
	AppPrivateKey string
	Charset       string
	SignType      string
}

const PrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA98lIxaX+Py6kBtf6TWtnnpLn1gGZV2s1CaXgNMvFw+Tej4no
vy9sxhwojr+cj74jFF6Jp3I5FJ/qCyW2htwc7VpE3dM5yLZlLvv7xeWzooVj/Oke
6HaRvdnZligonJHOVDxuMGzW8AikAo3NGpUm4ve7KOBs1dL8DZd2jxXECnk4hcgI
1YgEWz7JVfaQ/j/ArE1ADaIpALTg9HZVM7bJq5ZoGdsmCUd4C4vLbiLoh/rdK2vW
gkPDu9Zpd4Hk2UqWjDRE3GNkkIdaQmxKhUrtX1wuijcj6CE0IunueRBnBNUEpIjh
lFAzJ6dmqjR0Xlg0nh+Iiol6znTX1PcOGP6EBQIDAQABAoIBAQCUBzviV/g75rE3
JW/zMGcG5Nx7jRj+kJ1u1hnLcLEFBoWvWsQg80QYVlokbXQqq3xpftDdp+9R0vcP
EcipaHYflf3uR3IN5mksWH1hDIj0XpwNS3ebiLoooSzL99HLN4/74t4xL9R4MbFJ
lU0ixFgm37iAAxMB6rmJpSK++FHVrsiWxwvEoKbRfT7izThs4X06VGll9+wconRe
uYjAtm+S5hcc9/fELJm3bDwaWdN7ybMPji4yqtRiSJBaDd4p2NuUGDMOS+P1yPAH
c15x0ZUOQp/CVO6XYlsxsrvYz+CkJmJ6/sjWivxq6EM08mc04Lai2zVWFAYvMqph
EfJD+51BAoGBAP1SWq+GRsx+y2TDiU3Ig2lcb3ByLFKxSTidRBdLW/WmLZrHDz0o
eteAmwSk2gbzedb+PQP4FyS+W+fXwGuWoSdR/GjOcVH9GR0h39XmEuaMMi06BLcQ
yNIRHlGJVnqyMzlbw4phsRX5v6CdrpP3Spo7oLzCXWMbzvbKOvZOoIpnAoGBAPpn
8p2MesmvOrAATRJJjFwm2HN5w8ONyDzj68j7gVNie3aTP0WzDYEGxOsmdeGLG1fS
k1/6RAJlYlKHZDSUoummeRgq07f3QaIvfJ6k+QBxtvA6Z+D9u9ub54/jaMhO0d/9
VUUQYvB4DIHAbdbweCInyud/TAaE+giq7cO+WZKzAoGBAOcwWMkmL8kD0tZkShPt
8lie3qlt2ZuiZuO/S1xDD2sSPT6revHi1rGEknVbigub+09F+iN8MIr9G91sHxVR
hEPhZA22kt8zsM7QknqhHhDAVC7Ia3MzY0OsEdJyF7WkmnE3mS7a14Xpx4RrQ9+Q
acp2rsx2SkpgH7NFfyg5O/TzAoGBAKYH0khAYxHjS+hy6qdbeOOJJi+65uB82+3z
udzzVhaxz+cZTvSp+iQ5FsxMHhFEKQccKneS+xETpBPQjdKHU1XU+anai7MJEM22
6sxN2oQ+4et67nGyC6NbRjiTsmBOUr5PvQAkE1YaY0CNFMdVfnI3LEQ+lWwlM5wX
qbsGNWIdAoGALRjL+aXm+JhmDS03rToICKSjuJtMcCq3I4Mpk/hS3Kb/2VsDJIj7
oWXkRTSSkKBb7k8TYfb/2Yds1aPq3Qe2MPayJEUr3LdyQ1AVdIsjF7DUekQNcXf9
M/nGpj6f+aNXFudi14zzLwOptVttYL/iUHOtWC2PYfCUH6oipEy/wZI=
-----END RSA PRIVATE KEY-----
`

var DefaultAlipayClient = &AlipayClient{
	Url:           "https://openapi.alipay.com/gateway.do",
	AppID:         "2019051864987632",
	AppPrivateKey: PrivateKey,
	Charset:       "utf-8",
	SignType:      "RSA2",
}

type TokenResp struct {
	Rsp struct {
		AccessToken string `json:"access_token"`
		UserId      string `json:"user_id"`
	} `json:"alipay_system_oauth_token_response"`
}

type UserInfo struct {
	Rsp struct {
		UserId   string `json:"user_id"`
		Avatar   string `json:"avatar"`
		NickName string `json:"nick_name"`
	} `json:"alipay_user_info_share_response"`
}

func (c *AlipayClient) PostForm(data url.Values) ([]byte, error) {

	if err := c.sortAndSign(data); err != nil {
		return nil, err
	}

	resp, err := http.PostForm(c.Url, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		log.Println(string(bytes))

		return bytes, nil
	} else {
		return nil, errors.New("resp status not ok : " + strconv.Itoa(resp.StatusCode))
	}
}

func rsaSign(originData string, privateKey *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write([]byte(originData))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}

func (c *AlipayClient) NewRequest(method string) url.Values {
	values := make(url.Values)
	values.Set("app_id", c.AppID)
	values.Set("method", method)
	values.Set("charset", c.Charset)
	values.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	values.Set("version", "1.0")
	values.Set("sign_type", c.SignType)

	return values
}

func (c *AlipayClient) sortAndSign(values url.Values) error {
	p, rest := pem.Decode([]byte(c.AppPrivateKey))
	if p == nil {
		log.Println("private kek error ", string(rest))
		return errors.New("private key error")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		log.Println("x509 parse private key error :", err)
		return err
	}

	keys := make([]string, 0)
	for k, _ := range values {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var results []string
	for _, k := range keys {
		results = append(results, k+"="+values.Get(k))
	}

	sortFormData := strings.Join(results, "&")
	log.Println("sortFormData", sortFormData)

	sign, err := rsaSign(sortFormData, privateKey)
	if err != nil {
		return err
	}

	values.Set("sign", sign)
	return nil
}

// get Token and userId
func (c *AlipayClient) GetToken(code string) (*TokenResp, error) {
	req := c.NewRequest("alipay.system.oauth.token")
	req.Set("grant_type", "authorization_code")
	req.Set("code", code)

	bytes, err := c.PostForm(req)
	if err != nil {
		return nil, err
	}
	var tokenRsp TokenResp
	if err = json.Unmarshal(bytes, &tokenRsp); err != nil {
		return nil, err
	}
	return &tokenRsp, nil
}

func (c *AlipayClient) GetUserInfo(tokenResp *TokenResp) (string, error) {
	req := c.NewRequest("alipay.user.info.share")
	req.Set("auth_token", tokenResp.Rsp.AccessToken)

	bytes, err := c.PostForm(req)
	if err != nil {
		return "", err
	}

	log.Println(string(bytes))

	return string(bytes), nil
}

func AlipayLogin(code string) (id int, err error) {
	clinet := DefaultAlipayClient
	var tokenResp *TokenResp
	if tokenResp, err = clinet.GetToken(code); err != nil {
		return
	}

	var content string
	if content, err = clinet.GetUserInfo(tokenResp); err != nil {
		return
	}

	var info UserInfo
	if err = json.Unmarshal([]byte(content), &info); err != nil {
		return
	}

	if id, err = user.QueryUserIdByProfile(info.Rsp.UserId); err != nil {
		id, err = createUserByAlipay(&info, content)
		return
	}

	id, err = user.QueryUserIdByProfile(info.Rsp.UserId)
	return
}

func createUserByAlipay(info *UserInfo, content string) (int, error) {
	u := user.User{
		UserName:  info.Rsp.NickName,
		AvatarUrl: info.Rsp.Avatar,
	}
	if err := u.Create(); err != nil {
		return 0, err
	}

	subContent, _ := json.Marshal(info.Rsp)

	if err := user.CreateProfile(user.ALIPAY, u.Id, info.Rsp.UserId, string(subContent)); err != nil {
		return 0, err
	}

	return u.Id, nil
}
