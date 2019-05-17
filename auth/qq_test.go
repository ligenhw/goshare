package auth

// import (
// 	"log"
// 	"testing"
// )

// func TestGetAccessToken(t *testing.T) {
// 	token, err := GetAccessToken("66B63CB332B1A2E24C827181748E098D")
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
