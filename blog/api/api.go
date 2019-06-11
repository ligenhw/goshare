package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/ligenhw/goshare/auth"

	"github.com/ligenhw/goshare/session"

	"github.com/ligenhw/goshare/blog"
)

// /user
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	var result interface{}
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		bd := &blog.BlogDetail{
			Blog: blog.Blog{Id: id},
		}
		err = bd.QueryByID()
		if err != nil {
			return
		}
		result = bd
	} else {

		var limitN int
		if limitN, err = getQuery(r, "limit", 0); err != nil {
			return
		}

		var offsetN int
		if offsetN, err = getQuery(r, "offset", 0); err != nil {
			return
		}

		if userIDStr := r.URL.Query().Get("userId"); userIDStr == "" {
			if result, err = blog.GetAllBlogs(limitN, offsetN); err != nil {
				return
			}
		} else {
			var uid int
			if uid, err = strconv.Atoi(userIDStr); err != nil {
				return
			}
			if result, err = blog.GetArticleByUID(limitN, offsetN, uid); err != nil {
				return
			}
		}
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(result)
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

	_, err = auth.GetAuthUser(w, r)
	if err != nil {
		return
	}

	b.Time = time.Now()
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

	err = checkBlogOpPermission(w, r, b)
	if err != nil {
		return
	}

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

	err = checkBlogOpPermission(w, r, b)
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

func WithSession(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := session.Instance.SessionStart(w, r)
		if err != nil {
			log.Println(err)
			log.Println(session)
		}
		handler(w, r)
	}
}

func checkBlogOpPermission(w http.ResponseWriter, r *http.Request, blog blog.Blog) (err error) {

	var userID int
	userID, err = auth.GetAuthUser(w, r)

	err = blog.QueryById()
	if err != nil {
		return
	}

	if blog.UserId == userID {
		err = nil
	} else {
		err = errors.New(fmt.Sprintf("do not have blog op permission"))
	}

	return
}

func getQuery(r *http.Request, key string, defValue int) (value int, err error) {
	var valueStr string
	if valueStr = r.URL.Query().Get(key); valueStr == "" {
		value = defValue
		return
	}
	value, err = strconv.Atoi(valueStr)
	return
}

func init() {
	http.HandleFunc("/api/blog/", WithSession(BlogHandler))
	http.HandleFunc("/api/comments/", WithSession(CommentsHandler))
}
