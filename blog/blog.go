package blog

import (
	"time"

	"github.com/ligenhw/goshare/orm"

	"github.com/ligenhw/goshare/store"
)

type Blog struct {
	Id      int       `json:"id"`
	UserId  int       `json:"user_id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

var (
	db = store.Db
	o  = orm.NewOrm(store.Db)
)

func init() {
	orm.RegisterModel(new(Blog))
}

func (b *Blog) Create() (err error) {
	var id int64
	id, err = o.Insert(b)
	if err != nil {
		return
	}

	b.Id = int(id)
	return
}

// delete by id
func (b *Blog) Delete() (err error) {
	_, err = o.Delete(b)
	return
}

func (b *Blog) Update() (err error) {
	_, err = db.Exec("UPDATE blog SET title = ?, content = ? WHERE id = ?", b.Title, b.Content, b.Id)
	return
}

func (b *Blog) QueryById() (err error) {
	err = o.Read(b)
	return
}

func GetAllBlogs(limit, offset int) (blogs []*Blog, err error) {
	blogs = make([]*Blog, 0)
	qb := o.QueryTable(new(Blog))
	_, err = qb.LimitAndOffset(limit, offset).OrderBy("-time").All(&blogs)
	return
}

// GetArchieves get all article title
func GetArchieves() (blogs []*Blog, err error) {
	sel := []string{"id", "title", "time"}
	blogs = make([]*Blog, 0)
	qb := o.QueryTable(new(Blog))
	_, err = qb.OrderBy("-time").All(&blogs, sel...)
	return
}

func GetArticleByUID(limit, offset, uid int) (blogs []*Blog, err error) {
	blogs = make([]*Blog, 0)
	qb := o.QueryTable(new(Blog))
	_, err = qb.LimitAndOffset(limit, offset).Filter("user_id", uid).OrderBy("-time").All(&blogs)
	return
}
