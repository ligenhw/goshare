package blog

import (
	"testing"
	"time"

	"github.com/ligenhw/goshare/user"
)

func TestBlog(t *testing.T) {
	// set up
	u := user.User{UserName: "testuser_blog", Password: "testpass_blog"}
	if err := u.QueryByName(); err != nil {
		t.Log("no need delete testuser")
	} else if err := u.Delete(); err != nil {
		t.Error(err)
	}

	if err := u.Create(); err != nil {
		t.Error(err)
	}

	// create
	b := Blog{
		UserId:  u.Id,
		Title:   "testblog",
		Content: "testblog_content",
		Time:    time.Now(),
	}
	if err := b.Create(); err != nil {
		t.Error(err)
	}

	blogs, err := GetAllBlogs(100, 0)
	if err != nil {
		t.Error(err)
	}
	var find bool
	for _, item := range blogs {
		if item.Id == b.Id {
			find = true
			if item.Title != b.Title || item.Content != b.Content {
				t.Error("query all blogs not equals")
			}
		}
	}
	if !find {
		t.Error("do not have blog in getAllBlogs result.")
	}

	// update
	b.Title = "newTitle"
	if err = b.Update(); err != nil {
		t.Error(err)
	}

	// query
	b1 := Blog{
		Id: b.Id,
	}
	if err = b1.QueryById(); err != nil {
		t.Error(err)
	}
	if b1.Title != b.Title {
		t.Error("update title failed")
	}

	// delete
	if err = b.Delete(); err != nil {
		t.Error(err)
	}

	// clean
	if err = u.Delete(); err != nil {
		t.Error(err)
	}
}
