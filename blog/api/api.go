package api

import "net/http"

func HandleBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	}
}

func init() {
	http.HandleFunc("/blog", HandleBlog)
}
