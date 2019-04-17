package api

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/ligenhw/goshare/blog"
	"github.com/ligenhw/goshare/user"
)

type CommentsResp struct {
	Comments []*blog.Comment `json:"comments"`
	Users    []*user.User    `json:"users"`
}

// :blogId
func GetComment(w http.ResponseWriter, r *http.Request) (err error) {
	var blogId int
	blogId, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	var comments []*blog.Comment
	comments, err = blog.QueryByBlogId(blogId)
	if err != nil {
		return
	}

	var userIds []interface{}
	for _, comment := range comments {
		userIds = append(userIds, comment.UserId)
	}

	var users []*user.User
	users, err = user.QueryUserWithIds(userIds...)
	if err != nil {
		return
	}

	resp := CommentsResp{Comments: comments, Users: users}

	for _, c := range comments {
		log.Println(*c)
	}
	for _, u := range users {
		log.Println(*u)
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&resp)
	return
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Method)

	var err error
	switch r.Method {
	case http.MethodGet:
		err = GetComment(w, r)
		// case http.MethodPost:
		// 	err = Post(w, r)
		// case http.MethodDelete:
		// 	err = Delete(w, r)
		// case http.MethodPut:
		// 	err = Put(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
