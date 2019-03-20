package api

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/ligenhw/goshare/blog"
)

// /user
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	blogs, err := blog.GetAllBlogs()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(blogs)
	return
}

// /user
func Post(w http.ResponseWriter, r *http.Request) (err error) {
	// TODO: close the body
	decoder := json.NewDecoder(r.Body)
	b := blog.Blog{}
	err = decoder.Decode(&b)
	if err != nil {
		return
	}
	err = b.Create()
	return
}

// /user/:id
func Delete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	b := blog.Blog{Id: id}
	err = b.Delete()
	return
}

// /user
func Put(w http.ResponseWriter, r *http.Request) (err error) {
	decoder := json.NewDecoder(r.Body)
	b := blog.Blog{}
	err = decoder.Decode(&b)
	if err != nil {
		return
	}

	err = b.Update()
	return
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Method)

	var err error
	switch r.Method {
	case http.MethodGet:
		err = Get(w, r)
	case http.MethodPost:
		err = Post(w, r)
	case http.MethodDelete:
		err = Delete(w, r)
	case http.MethodPut:
		err = Put(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/blog/", BlogHandler)
}
