package blog

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetAllBlogs(t *testing.T) {
	blogs, _ := GetAllBlogs()
	for _, blog := range blogs {
		t.Log(blog)
	}
}

func TestCreate(t *testing.T) {
	b := Blog{User_Id: 1, Title: "testT1", Content: "testC1"}
	t.Log("Create : ", b.Create())

	b.Title = "newttttttitle"
	b.Id = 7
	t.Log("Update : ", b.Update())

	t.Log("Delete : ", b.Delete())

}

func TestCreateFromFile(t *testing.T) {
	path := "../script/testdata/"
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		t.Log("info ", info.Name())
		name := info.Name()
		if content, err := ioutil.ReadFile(path + name); err == nil {
			b := Blog{User_Id: 1, Title: strings.Split(name, ".")[0], Content: string(content)}
			b.Create()
		}
	}
}
