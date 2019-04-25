package api

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

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/user/", UserHandler)
	writer = httptest.NewRecorder()
}

func TestPostApi(t *testing.T) {
	username := "testuser_api"
	u := &user.User{UserName: username, Password: "testpass_api"}
	bs, _ := json.Marshal(u)
	request, _ := http.NewRequest("POST", "/api/user/", bytes.NewReader(bs))

	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v %s", writer.Code, writer.Body)
	}

	// clean up
	u = &user.User{UserName: username}
	u.QueryByName()
	u.Delete()
}
