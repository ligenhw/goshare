package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
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
	token_url     = "https://github.com/login/oauth/access_token"
	client_id     = "6f8ed9e6e9fc7b14cbc2"
	client_secret = "d0b7a458886df8b4f6ed4f9a3684e88f67db99b1"
	CONTENT_TYPE  = "application/json"
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
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

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

func getGhUserInfo(token string) (err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", user_url, nil)
	req.Header.Add("Authorization", "token "+token)

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return GET_USER_INFO_ERR
	} else {
		defer resp.Body.Close()

		var bs []byte
		bs, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		log.Println("get user info from gh : ")
		log.Println(string(bs))
		return
	}
}
