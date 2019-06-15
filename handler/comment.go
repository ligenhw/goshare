package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ligenhw/goshare/handler/context"

	"github.com/gorilla/mux"

	"github.com/ligenhw/goshare/blog"
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
	userID := *context.UserID(r)

	defer r.Body.Close()
	c := CreateCommentsReq{}
	var err error

	if err = json.NewDecoder(r.Body).Decode(&c); err != nil {
		handleError(err, w)
		return
	}

	if c.ParentCommentID != nil {
		err = blog.CreateReply(c.BlogID, userID, *c.ParentCommentID, c.ReplyTo, c.Content)
	} else {
		_, err = blog.CreateComment(c.BlogID, userID, c.Content)
	}
	handleError(err, w)

	return
}

// GetComment get comment list by blogId
func GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogID, err := strconv.Atoi(vars["blogId"])
	if err != nil {
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
