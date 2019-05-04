package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ligenhw/goshare/configration"
	"github.com/ligenhw/goshare/user"
)

type GhReq struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type GhResp struct {
	AccessToken string `json:"access_token"`
}

const (
	token_url    = "https://github.com/login/oauth/access_token"
	CONTENT_TYPE = "application/json"
)

var (
	client_id     = configration.Conf.ClientId
	client_secret = configration.Conf.ClientSecret
)

var (
	GET_TOKEN_ERR = errors.New("get token failed")
)

func getGhToken(code string) (token string, err error) {
	reqBody := GhReq{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Code:         code,
	}

	var bs []byte
	bs, err = json.Marshal(reqBody)
	if err != nil {
		return
	}

	var req *http.Request
	req, err = http.NewRequest("POST", token_url, bytes.NewReader(bs))
	req.Header.Set("Content-Type", CONTENT_TYPE)
	req.Header.Set("Accept", CONTENT_TYPE)

	var resp *http.Response
	log.Println("request to : ", token_url, "start")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	log.Println("request to : ", token_url, "end")

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		var respBytes []byte
		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		log.Println(string(respBytes))

		var ghresp GhResp
		err = json.Unmarshal(respBytes, &ghresp)
		if err != nil {
			return
		}

		log.Println("get token : ", ghresp.AccessToken)
		return ghresp.AccessToken, nil
	} else {
		log.Println(token_url, "response code : ", resp.StatusCode)
		err = GET_TOKEN_ERR
		return
	}
}

const (
	user_url = "https://api.github.com/user"
)

var (
	GET_USER_INFO_ERR = errors.New("get user info failed")
)

type GhUserInfo struct {
	Login     string `json:"login"`
	Id        int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Email     string `json:"email"`
	Location  string `json:"location"`
}

func getGhUserInfo(token string) (info *GhUserInfo, err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", user_url, nil)
	req.Header.Add("Authorization", "token "+token)

	var resp *http.Response
	log.Println("request to : ", user_url, "start")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	log.Println("request to : ", user_url, " end")

	if resp.StatusCode != http.StatusOK {
		return nil, GET_USER_INFO_ERR
	} else {
		defer resp.Body.Close()

		var bs []byte
		bs, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		log.Println("get user info from gh : ")
		log.Println(string(bs))

		info = new(GhUserInfo)
		err = json.Unmarshal(bs, &info)
		if err != nil {
			return nil, err
		}
		return
	}
}

func GhLogin(code string) (id int, err error) {
	var token string
	token, err = getGhToken(code)
	if err != nil {
		return
	}

	var info *GhUserInfo
	info, err = getGhUserInfo(token)
	if err != nil {
		return
	}

	u := user.User{
		UserName:  info.Login,
		AvatarUrl: info.AvatarUrl,
		Time:      time.Now(),
	}

	err = u.QueryByName()
	if err != nil {
		log.Println(err)
		err = u.Create()
		if err != nil {
			return
		}
	}

	id = u.Id
	return
}
