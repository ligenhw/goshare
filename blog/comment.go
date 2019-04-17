package blog

import (
	"time"

	"github.com/ligenhw/goshare/orm"
	"github.com/ligenhw/goshare/store"
)

type Comment struct {
	Id      int       `json:"id"`
	BlogId  int       `json:"blogId"`
	UserId  int       `json:"userId"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

var (
	o = orm.NewOrm(store.Db)
)

func init() {
	orm.RegisterModel(new(Comment))
}

func CreateComment(blogId, userId int, content string) (err error) {
	b := &Comment{
		BlogId:  blogId,
		UserId:  userId,
		Content: content,
	}

	_, err = o.Insert(b)
	return
}

func QueryByBlogId(blogId int) (comments []*Comment, err error) {
	comments = make([]*Comment, 0)
	qs := o.QueryTable(new(Comment))
	_, err = qs.Filter("blog_id", blogId).All(&comments)
	return
}
