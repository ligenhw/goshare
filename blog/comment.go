package blog

import (
	"time"

	"github.com/ligenhw/goshare/orm"
)

type Comment struct {
	Id       int       `json:"id"`
	BlogId   int       `json:"blogId"`
	UserId   int       `json:"userId"`
	ParentId *int      `json:"parentId" orm:"null"`
	ReplyTo  *int      `json:"replyTo" orm:"null"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

func CreateComment(blogId, userId int, content string) (err error) {
	b := &Comment{
		BlogId:  blogId,
		UserId:  userId,
		Content: content,
		Time:    time.Now(),
	}

	_, err = o.Insert(b)
	return
}

func QueryByBlogId(blogId int) (comments []*Comment, err error) {
	comments = make([]*Comment, 0)
	qs := o.QueryTable(new(Comment))
	_, err = qs.Filter("blog_id", blogId).OrderBy("-time").All(&comments)
	return
}
