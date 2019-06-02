package api

import (
	"encoding/json"
	"net/http"

	"github.com/ligenhw/goshare/blog"
)

func TagHandler(w http.ResponseWriter, r *http.Request) {
	if tags, err := blog.GetTags(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		bytes, _ := json.Marshal(tags)
		w.Header().Set("content-type", "application/json")
		w.Write(bytes)
	}
}

func init() {
	http.HandleFunc("/api/tag/", TagHandler)
}
