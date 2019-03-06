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
			if _, err := w.Write(b); err != nil {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			log.Panicln(err)
		}
	}
}

func init() {
	http.HandleFunc("/user", UserHandler)
}
