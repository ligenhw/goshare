package api

import (
	"encoding/json"
	"net/http"

	"github.com/ligenhw/goshare/session"
)

func handleSession(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(session.Instance.GetProvider())
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func init() {
	http.HandleFunc("/api/admin/sessions/", handleSession)
}
