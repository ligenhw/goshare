package api

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/session"

	"github.com/ligenhw/goshare/blog"
	"github.com/ligenhw/goshare/user"
)

type CommentsResp struct {
	Comments []*blog.Comment `json:"comments"`
	Users    []*user.User    `json:"users"`
}

type CommentReq struct {
	Content string `json:"content"`
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

	users := make([]*user.User, 0)
	if len(userIds) > 0 {
		users, err = user.QueryUserWithIds(userIds...)
		if err != nil {
			return
		}
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

// :blogId
func CreateComment(w http.ResponseWriter, r *http.Request) (err error) {
	var blogId int
	blogId, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	globalSession := session.Instance

	var session session.Store
	session, err = globalSession.SessionStart(w, r)
	if err != nil {
		return
	}

	var userID int
	userID, err = auth.Auth(session)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	c := CommentReq{}
	err = decoder.Decode(&c)
	if err != nil {
		return
	}

	err = blog.CreateComment(blogId, userID, c.Content)
	return
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Method)

	var err error
	switch r.Method {
	case http.MethodGet:
		err = GetComment(w, r)
	case http.MethodPost:
		err = CreateComment(w, r)
		// case http.MethodDelete:
		// 	err = Delete(w, r)
		// case http.MethodPut:
		// 	err = Put(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
