package auth

import "testing"

func TestGhAuth(t *testing.T) {
	code := "eff71b85fea38d7e17fe"
	token, err := getGhToken(code)
	if err != nil {
		t.Fatal(err)
	}

	err = getGhUserInfo(token)
	if err != nil {
		t.Fatal(err)
	}
}
