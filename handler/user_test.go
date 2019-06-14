package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ligenhw/goshare/user"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

var smux *http.ServeMux
var writer *httptest.ResponseRecorder

func setUp() {
	smux = http.NewServeMux()
	smux.HandleFunc("/api/user", CreateUser)
	writer = httptest.NewRecorder()
}

func TestPostApi(t *testing.T) {
	username := "testuser_api"
	u := &user.User{UserName: username, Password: "testpass_api"}
	bs, _ := json.Marshal(u)
	request, _ := http.NewRequest("POST", "/api/user", bytes.NewReader(bs))

	smux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v %s", writer.Code, writer.Body)
	}

	// clean up
	u = &user.User{UserName: username}
	u.QueryByName()
	u.Delete()
}
