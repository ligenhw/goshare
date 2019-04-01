package main

import (
	"net/http"

	_ "github.com/ligenhw/goshare/auth/api"
	_ "github.com/ligenhw/goshare/blog/api"
	_ "github.com/ligenhw/goshare/health/api"
	"github.com/ligenhw/goshare/session"
	_ "github.com/ligenhw/goshare/session/api"
	_ "github.com/ligenhw/goshare/user/api"
	"github.com/ligenhw/goshare/version"
)

var globalSession *session.Manager

func init() {
	globalSession, _ = session.NewManager("mem")
	go globalSession.GC()
}

func main() {
	p("Go share", version.Version, "started at", config.Address)
	http.ListenAndServe(config.Address, nil)
}
