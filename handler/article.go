package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ligenhw/goshare/blog"
)

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

// GetArticles : get all articles
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
