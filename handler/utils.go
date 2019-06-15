package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ligenhw/goshare/blog"
)

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func checkOwnerOpPerm(w http.ResponseWriter, r *http.Request, blog blog.Blog, userID int) (err error) {

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
