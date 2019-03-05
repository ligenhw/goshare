package api

import (
	"errors"
	"net/http"

	"github.com/ligenhw/goshare/health"
)

var updater = health.NewStatusUpdater()

func DownHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		updater.Update(errors.New("manual Check"))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func UpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		updater.Update(nil)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func init() {
	health.Register("manual_http_status", updater)
	http.HandleFunc("/debug/health/down", DownHandler)
	http.HandleFunc("/debug/health/up", UpHandler)
}
