package main

import (
	"fmt"
	"net/http"

	_ "github.com/ligenhw/goshare/health/api"
	_ "github.com/ligenhw/goshare/user/api"
	"github.com/ligenhw/goshare/version"
)

func main() {
	fmt.Println("version : " + version.Version)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":5001", nil)
}
