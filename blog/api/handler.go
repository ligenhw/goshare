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
	Comments []*blog.CommentWithChild `json:"comments"`
	Users    []*user.User             `json:"users"`
}

type CreateCommentsReq struct {
	BlogId          int    `json:"blogId"`
	ParentCommentId *int   `json:"parentCommentId"`
	ReplyTo         int    `json:"replyTo"`
	Content         string `json:"content"`
}

// :blogId
func GetComment(w http.ResponseWriter, r *http.Request) (err error) {
	var blogId int
	blogId, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	var comments []*blog.CommentWithChild
	comments, err = blog.QueryCommentsByBlogId(blogId)
	if err != nil {
		return
	}

	userIdsMap := make(map[int]bool, 0)
	for _, comment := range comments {
		userIdsMap[comment.ParentUserId] = true
		for _, sub := range comment.SubComments {
			if sub == nil {
				continue
			}
			userIdsMap[sub.UserId] = true
		}
	}

	var users []*user.User
	userIds := make([]interface{}, 0)
	if len(userIdsMap) > 0 {
		for k, _ := range userIdsMap {
			userIds = append(userIds, k)
		}
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
	err = json.NewEncoder(w).Encode(&resp)
	return
}

// :blogId
func CreateComment(w http.ResponseWriter, r *http.Request) (err error) {
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

	c := CreateCommentsReq{}
	err = decoder.Decode(&c)
	if err != nil {
		return
	}

	if c.ParentCommentId != nil {
		err = blog.CreateReply(c.BlogId, userID, *c.ParentCommentId, c.ReplyTo, c.Content)
	} else {
		_, err = blog.CreateComment(c.BlogId, userID, c.Content)
	}

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
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
