package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ligenhw/goshare/blog"
	"github.com/ligenhw/goshare/user"
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
eg: db -d ../../script/testdata/
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

	if flag.NArg() > 0 {
		for _, file := range flag.Args() {
			if err := scanFile(file); err != nil {
				panic(err)
			}
		}
	}
}

func scanDir(path string) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		name := info.Name()
		err := scanFile(filepath.Join(path, name))
		if err != nil {
			panic(err)
		}
	}
}

func getOrCreateUser() int {
	users, err := user.GetAllUser()
	if err != nil {
		panic(err)
	}

	if len(users) > 0 {
		return users[0].Id
	} else {
		u := &user.User{
			UserName: "test",
			Password: "pass",
		}
		if err := u.Create(); err != nil {
			panic(err)
		}

		return u.Id
	}
}

func scanFile(path string) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(path); err == nil {
		name := filepath.Base(path)
		log.Println("scan file : ", name)
		b := blog.Blog{UserId: getOrCreateUser(), Title: strings.Split(name, ".")[0], Content: string(content)}
		err = b.Create()
	}

	return
}
