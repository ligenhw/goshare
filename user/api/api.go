package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ligenhw/goshare/user"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users := user.GetAllUser()
		if b, err := json.Marshal(users); err == nil {
			w.Write(b)
		} else {
			log.Panicln(err)
		}
	}
}

func init() {
	http.HandleFunc("/user", UserHandler)
}
