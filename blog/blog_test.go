package blog

import (
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

func TestBlogDetails(t *testing.T) {
	bd := BlogDetail{
		Blog: Blog{
			Id: 55,
		},
	}
	t.Log(bd.QueryByID())
	t.Log(bd)
}
