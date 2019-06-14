package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/blog"
	"github.com/ligenhw/goshare/session"
	"github.com/ligenhw/goshare/user"
)

// CommentsResp body
type CommentsResp struct {
	Comments []*blog.CommentWithChild `json:"comments"`
	Users    []*user.User             `json:"users"`
}

// CreateCommentsReq body
type CreateCommentsReq struct {
	BlogID          int    `json:"blogId"`
	ParentCommentID *int   `json:"parentCommentId"`
	ReplyTo         int    `json:"replyTo"`
	Content         string `json:"content"`
}

// CreateComment create a comment or reply
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var err error
	globalSession := session.Instance
	var session session.Store
	if session, err = globalSession.SessionStart(w, r); err != nil {
		handleError(err, w)
		return
	}

	var userID int
	if userID, err = auth.Auth(session); err != nil {
		handleError(err, w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	c := CreateCommentsReq{}
	if err = decoder.Decode(&c); err != nil {
		handleError(err, w)
		return
	}

	if c.ParentCommentID != nil {
		err = blog.CreateReply(c.BlogID, userID, *c.ParentCommentID, c.ReplyTo, c.Content)
	} else {
		_, err = blog.CreateComment(c.BlogID, userID, c.Content)
	}

	return
}

// GetComment get comment list by blogId
func GetComment(w http.ResponseWriter, r *http.Request) {
	var err error
	var blogID int
	vars := mux.Vars(r)
	if blogID, err = strconv.Atoi(vars["blogId"]); err != nil {
		handleError(err, w)
		return
	}

	var comments []*blog.CommentWithChild
	if comments, err = blog.QueryCommentsByBlogId(blogID); err != nil {
		handleError(err, w)
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
		for k := range userIdsMap {
			userIds = append(userIds, k)
		}
		if users, err = user.QueryUserWithIds(userIds...); err != nil {
			handleError(err, w)
			return
		}
	}

	resp := CommentsResp{Comments: comments, Users: users}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	handleError(err, w)
	return
}
