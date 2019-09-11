package blog

import (
	"strconv"
	"testing"
	"time"

	"github.com/ligenhw/goshare/user"
)

func TestComment(t *testing.T) {
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

	u2 := user.User{UserName: "testuser_blog2", Password: "testpass_blog2"}
	if err := u2.QueryByName(); err != nil {
		t.Log("no need delete testuser")
	} else if err := u2.Delete(); err != nil {
		t.Error(err)
	}

	if err := u2.Create(); err != nil {
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

	c1, err := CreateComment(b.Id, u.Id, "test comment1")
	if err != nil {
		t.Error(err)
	}
	c2, err := CreateComment(b.Id, u.Id, "test comment2")
	if err != nil {
		t.Error(err)
	}
	if _, err := CreateComment(b.Id, u.Id, "test Comment3🐟"); err != nil {
		t.Fatal(err)
	}

	CreateReply(b.Id, u2.Id, int(c1), u.Id, "reply to c1")
	CreateReply(b.Id, u.Id, int(c1), u2.Id, "i am receive reply from u")

	CreateReply(b.Id, u2.Id, int(c2), u.Id, "reply to c2")
	CreateReply(b.Id, u2.Id, int(c2), u.Id, "reply to c2 twice")

	comments := queryCommentsByBlogId(t, b.Id)

	if len(comments) != 3 {
		t.Error("comments test error! err count : ", len(comments))
	}

	// delete
	if err := b.Delete(); err != nil {
		t.Error(err)
	}

	// clean
	if err = u.Delete(); err != nil {
		t.Error(err)
	}

	if err = u2.Delete(); err != nil {
		t.Error(err)
	}
}

func queryCommentsByBlogId(t *testing.T, blogId int) (comments []*CommentWithChild) {
	var err error
	if comments, err = QueryCommentsByBlogId(blogId); err != nil {
		t.Error(err)
	}

	for i, c := range comments {
		t.Log("Comment : " + strconv.Itoa(i) + " ----------------------")
		t.Log(*c)
		if c.SubComments == nil {
			continue
		}
		for _, sub := range c.SubComments {
			t.Logf("    %v", sub)
		}
		t.Log("--------------------------------------------")
	}

	return
}
