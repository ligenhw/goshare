package auth

// import (
// 	"log"
// 	"testing"
// )

// func TestGetAccessToken(t *testing.T) {
// 	token, err := GetAccessToken("BC03BA88AC21CFCD167B381B837F7C78")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Println("token --- ", token)

// 	openid, err := GetOpenID(token)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Println("openid --- ", openid)

// 	info, err := GetQQUserInfo(token, openid)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Println(info.Nickname)
// }
