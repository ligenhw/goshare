package main

import (
	"net/http"

	_ "github.com/ligenhw/goshare/blog/api"
	_ "github.com/ligenhw/goshare/health/api"
	_ "github.com/ligenhw/goshare/session/api"
	_ "github.com/ligenhw/goshare/user/api"
	"github.com/ligenhw/goshare/version"
)

func main() {
	p("Go share", version.Version, "started at", config.Address)
	http.ListenAndServe(config.Address, nil)
}
