package main

import (
	"log"
	"net/http"

	"github.com/ligenhw/goshare/handler/middleware"

	"github.com/gorilla/mux"
	"github.com/ligenhw/goshare/configration"
	"github.com/ligenhw/goshare/handler"
	"github.com/ligenhw/goshare/session"
	"github.com/ligenhw/goshare/version"
)

var (
	globalSession *session.Manager
)

func init() {
	log.SetFlags(log.Flags() | log.Llongfile)

	globalSession, _ = session.NewManager("mem")
	go globalSession.GC()
}

var (
	u = middleware.CheckUser
)

func main() {
	p("Go share", version.Version, "started at", configration.Conf.Address)

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/api/user", u(handler.GetUser)).Methods("GET")
	r.HandleFunc("/api/user", handler.CreateUser).Methods("POST")
	r.HandleFunc("/api/login", handler.Login).Methods("POST")
	r.HandleFunc("/api/logout", handler.Logout).Methods("POST")

	r.HandleFunc("/api/article", handler.GetArticles).Methods("GET")
	r.HandleFunc("/api/article", u(handler.CreateArticle)).Methods("POST")
	r.HandleFunc("/api/article", u(handler.UpdateArticle)).Methods("PUT")
	r.HandleFunc("/api/article/{id}", handler.GetArticleByID).Methods("GET")
	r.HandleFunc("/api/article/{id}", u(handler.DeleteArticle)).Methods("DELETE")
	r.HandleFunc("/api/archives", handler.GetArchives).Methods("GET")

	r.HandleFunc("/api/comment/{blogId}", handler.GetComment).Methods("GET")
	r.HandleFunc("/api/comment", u(handler.CreateComment)).Methods("POST")

	r.HandleFunc("/api/tag", handler.TagHandler)

	r.HandleFunc("/api/ghlogin", handler.GhLogin).Methods("POST")
	r.HandleFunc("/api/qqlogin", handler.QqLogin).Methods("POST")
	r.HandleFunc("/api/alipaylogin", handler.AlipayLogin).Methods("POST")

	http.Handle("/", r)

	if err := http.ListenAndServe(configration.Conf.Address, nil); err != nil {
		panic(err)
	}
}
