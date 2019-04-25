package blog

import (
	"time"

	"github.com/ligenhw/goshare/orm"

	"github.com/ligenhw/goshare/store"
)

type Blog struct {
	Id      int       `json:"id"`
	User_Id int       `json:"user_id"`
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

func GetAllBlogs() (blogs []*Blog, err error) {
	rows, err := db.Query("SELECT id, user_id, title, content, time FROM blog ORDER BY time DESC")
	if err != nil {
		return
	}

	for rows.Next() {
		b := Blog{}
		err = rows.Scan(&b.Id, &b.User_Id, &b.Title, &b.Content, &b.Time)
		if err != nil {
			return
		}
		blogs = append(blogs, &b)
	}
	return
}
