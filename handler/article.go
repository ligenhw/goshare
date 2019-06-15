package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ligenhw/goshare/blog"
	"github.com/ligenhw/goshare/handler/context"
)

// TODO ; how to handle the error ? goto end

// GetArticleByID : get single article and user info
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleError(err, w)
		return
	}

	bd := &blog.BlogDetail{
		Blog: blog.Blog{Id: id},
	}
	if err = bd.QueryByID(); err != nil {
		handleError(err, w)
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
	b := blog.Blog{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		handleError(err, w)
		return
	}
	b.UserId = *context.UserID(r)

	b.Time = time.Now()
	err = b.Create()
	handleError(err, w)
	return
}

// UpdateArticle  update a single article
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	userID := *context.UserID(r)

	b := blog.Blog{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		handleError(err, w)
		return
	}
	if err = checkOwnerOpPerm(w, r, b, userID); err != nil {
		handleError(err, w)
		return
	}

	err = b.Update()
	handleError(err, w)
	return
}

// DeleteArticle  delete a single article, check owner permission
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	userID := *context.UserID(r)

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleError(err, w)
		return
	}
	b := blog.Blog{Id: id}
	if err = checkOwnerOpPerm(w, r, b, userID); err != nil {
		handleError(err, w)
		return
	}

	err = b.Delete()
	handleError(err, w)
	return
}
