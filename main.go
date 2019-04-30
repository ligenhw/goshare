package main

import (
	"log"
	"net/http"

	_ "github.com/ligenhw/goshare/auth/api"
	_ "github.com/ligenhw/goshare/blog/api"
	"github.com/ligenhw/goshare/configration"
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
	http.ListenAndServe(configration.Conf.Address, nil)
}
