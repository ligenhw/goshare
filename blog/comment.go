package blog

import (
	"time"

	"github.com/ligenhw/goshare/orm"
	"github.com/ligenhw/goshare/store"
)

type Comment struct {
	Id      int
	BlogId  int
	UserId  int
	Content string
	Time    time.Time `orm:"-"`
}

var (
	o = orm.NewOrm(store.Db)
)

func init() {
	orm.RegisterModel(new(Comment))
}

func Create(blogId, userId int, content string) (err error) {
	b := &Comment{
		BlogId:  blogId,
		UserId:  userId,
		Content: content,
	}

	_, err = o.Insert(b)
	return
}

func QueryByBlogId(blogId int) (err error) {
	return
}
