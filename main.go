package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/ligenhw/goshare/auth/api"
	_ "github.com/ligenhw/goshare/blog/api"
	"github.com/ligenhw/goshare/configration"
	"github.com/ligenhw/goshare/handler"
	_ "github.com/ligenhw/goshare/health/api"
	"github.com/ligenhw/goshare/session"
	_ "github.com/ligenhw/goshare/session/api"
	_ "github.com/ligenhw/goshare/user/api"
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

func main() {
	p("Go share", version.Version, "started at", configration.Conf.Address)

	r := mux.NewRouter()

	r.HandleFunc("/api/article", handler.GetArticles).Methods("GET")
	r.HandleFunc("/api/article/{id}", handler.GetArticleByID).Methods("GET")

	r.HandleFunc("/api/tag", handler.TagHandler)

	http.Handle("/", r)

	err := http.ListenAndServe(configration.Conf.Address, nil)

	if err != nil {
		panic(err)
	}
}
