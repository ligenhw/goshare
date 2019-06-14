package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/blog"
)

// TODO ; how to handle the error ? goto end

// GetArticleByID : get single article and user info
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id int
	var err error

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return
	}

	bd := &blog.BlogDetail{
		Blog: blog.Blog{Id: id},
	}
	if err = bd.QueryByID(); err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(bd)
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

// GetArticles : get all articles query param limit , offset , userId
func GetArticles(w http.ResponseWriter, r *http.Request) {
	var limitN int
	var err error

	if limitN, err = getQuery(r, "limit", 0); err != nil {
		return
	}

	var offsetN int
	if offsetN, err = getQuery(r, "offset", 0); err != nil {
		return
	}

	var result []*blog.Blog

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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	return
}

// CreateArticle create a new article .
// permission: login
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// TODO: close the body
	var err error
	b := blog.Blog{}
	if err = json.NewDecoder(r.Body).Decode(&b); err != nil {
		return
	}

	b.Time = time.Now()
	err = b.Create()
	return
}

// UpdateArticle  update a single article
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	b := blog.Blog{}
	if err = json.NewDecoder(r.Body).Decode(&b); err != nil {
		return
	}

	if err = checkBlogOpPermission(w, r, b); err != nil {
		return
	}

	err = b.Update()
	return
}

func checkBlogOpPermission(w http.ResponseWriter, r *http.Request, blog blog.Blog) (err error) {

	var userID int
	userID, err = auth.GetAuthUser(w, r)

	if err = blog.QueryById(); err != nil {
		return
	}

	if blog.UserId == userID {
		err = nil
	} else {
		err = fmt.Errorf("do not have blog op permission")
	}

	return
}

// DeleteArticle  delete a single article, check owner permission
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int
	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return
	}

	b := blog.Blog{Id: id}

	if err = checkBlogOpPermission(w, r, b); err != nil {
		return
	}

	err = b.Delete()
	return
}
