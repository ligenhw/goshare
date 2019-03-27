package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ligenhw/goshare/blog"
)

type args string

func (a args) String() string {
	return string(a)
}

func (a *args) Set(value string) error {
	*a = args(value)
	return nil
}

var (
	dir         args
	showVersion bool
)

func init() {
	flag.Var(&dir, "d", "scan the md file in the directionary")
	flag.BoolVar(&showVersion, "version", false, "show version and exits")

	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [files...]\n", os.Args[0])
	fmt.Fprint(os.Stderr, `
Scan the md files and Insert to blog table
maybe insert as user 1.
`)
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Println("version 1.0")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if showVersion {
		printVersion()
		return
	}

	if dir != "" {
		scanDir(dir.String())
	}

}

func scanDir(path string) {
	// path := "../../script/testdata/"
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		log.Println("info ", info.Name())
		name := info.Name()
		scanFile(path, name)
	}
}

func scanFile(path, name string) (err error) {
	if content, err := ioutil.ReadFile(path + name); err == nil {
		b := blog.Blog{User_Id: 1, Title: strings.Split(name, ".")[0], Content: string(content)}
		b.Create()
	}

	return
}
