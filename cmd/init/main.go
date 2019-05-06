package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ligenhw/goshare/store"
)

var (
	drop bool
	file string
)

func init() {
	flag.BoolVar(&drop, "drop", false, "drop goshare db")
	flag.StringVar(&file, "file", "../../script/schema.sql", "db schema sql file")

	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [OPTION]... [FILE]\n", os.Args[0])
	fmt.Fprint(os.Stderr, `
Init the goshare db schema.

ENVIRONMENT
  export DSN="gen:1234@tcp(192.168.199.231)/mysql?charset=utf8&parseTime=true"
eg: init --drop --file schema.sql
`)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	err := store.InitDb(file, drop)
	if err != nil {
		log.Println("db init success")
	}
}
