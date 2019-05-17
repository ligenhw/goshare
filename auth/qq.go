package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const (
	APP_ID        = "101576375"
	APP_SECRET    = ""
	REDIRECT_URI  = "http://bestlang.cn/login"
	TOKEN_URL     = "https://graph.qq.com/oauth2.0/token?"
	USER_INFO_URL = "https://graph.qq.com/user/get_user_info?"
)

func GetTokenUrl(code string) string {
	val := make(url.Values)
	val.Add("grant_type", "authorization_code")
	val.Add("client_id", APP_ID)
	val.Add("client_secret", APP_SECRET)
	val.Add("code", code)
	val.Add("redirect_uri", REDIRECT_URI)
	p := val.Encode()

	return TOKEN_URL + p
}

func GetAccessToken(code string) (token string, err error) {
	var resp *http.Response

	log.Println("qq GetAccessToken code :", code)
	token_url := GetTokenUrl(code)
	log.Println(token_url)
	resp, err = http.Get(token_url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	log.Println("qq GetAccessToken resp :", resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		var respBytes []byte
		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		log.Println(string(respBytes))
		var values url.Values
		if values, err = url.ParseQuery(string(respBytes)); err != nil {
			log.Println(err)
			return
		}

		token = values.Get("access_token")
	} else {
		err = GET_TOKEN_ERR
	}
	return
}

func GetOpenID(accessToken string) (openid string, err error) {
	url := "https://graph.qq.com/oauth2.0/me?access_token=" + accessToken
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	if http.StatusOK == resp.StatusCode {
		var bytes []byte
		if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
			log.Println(err)
			return
		}

		log.Println(string(bytes))
		contentRe := regexp.MustCompile(`^callback\(\s+(\S+)\s+\);`)
		if content := contentRe.FindStringSubmatch(string(bytes)); len(content) == 2 {
			log.Println(content)
			jsonContent := content[1]
			var callback struct {
				ClientId string `json:"client_id"`
				OpenId   string `json:"openid"`
			}
			if err = json.Unmarshal([]byte(jsonContent), &callback); err != nil {
				log.Println(err)
				return
			}

			log.Println(callback)
			openid = callback.OpenId
		} else {
			log.Println(content)
			err = errors.New("response formate error")
		}
	} else {
		err = errors.New("response status not ok")
	}

	return
}

type QQUserInfo struct {
	Nickname string `json:"nickname"`
}

func GetQQUserInfoUrl(accessToken, openid string) string {
	val := make(url.Values)
	val.Add("access_token", accessToken)
	val.Add("oauth_consumer_key", APP_ID)
	val.Add("openid", openid)
	p := val.Encode()

	return USER_INFO_URL + p
}

func GetQQUserInfo(accessToken, openid string) (info *QQUserInfo, err error) {
	url := GetQQUserInfoUrl(accessToken, openid)

	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	if http.StatusOK == resp.StatusCode {
		var bytes []byte
		if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
			log.Println(err)
			return
		}

		log.Println(string(bytes))

		info = new(QQUserInfo)
		if err = json.Unmarshal(bytes, info); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return
}
