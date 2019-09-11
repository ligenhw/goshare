package blog

import (
	"testing"
	"time"

	"github.com/ligenhw/goshare/user"
)

func TestBlogDetail(t *testing.T) {
	u := user.User{
		UserName: "testuser_bd",
		Password: "testpass_db",
	}
	if err := u.QueryByName(); err != nil {
		if err := u.Create(); err != nil {
			t.Fatal(err)
		}
	}

	b := Blog{
		UserId:  u.Id,
		Title:   "title",
		Content: "content",
		Time:    time.Now(),
	}
	if err := b.Create(); err != nil {
		t.Error(err)
	}

	bd := BlogDetail{
		Blog: Blog{
			Id: b.Id,
		},
	}
	if err := bd.QueryByID(); err != nil {
		t.Error(err)
	}
	if bd.User.Id != u.Id || bd.Title != b.Title {
		t.Error("blog detail query failed.")
	}

	if err := u.Delete(); err != nil {
		t.Fatal(err)
	}

}
