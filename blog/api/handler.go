package api

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/ligenhw/goshare/blog"
)

// :blogId
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	blogId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	var comments []blog.Comment
	comments, err = blog.QueryByBlogId(blogId)
	if err != nil {
		return
	}

	var userIds []int
	for _, comment := range comments {
		userIds = append(userIds, comment.UserId)
	}

	

}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
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
