package blog

import (
	"time"

	"github.com/ligenhw/goshare/store"
)

type Blog struct {
	Id      int       `json:"id"`
	User_Id int       `json:"user_id"`
	Title  string    `json:"title"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

type Comment struct {
	Id      string
	BlogId  string
	UserId  string
	Content string
	Time    time.Time
}

var db = store.Db

func (b *Blog) Create() (err error) {
	_, err = db.Exec("INSERT INTO blog (user_id, title, content) VALUES (?, ?, ?)", b.User_Id, b.Title, b.Content)
	return
}

// delete by id
func (b *Blog) Delete() (err error) {
	_, err = db.Exec("DELETE FROM blog WHERE id = ?", b.Id)
	return
}

func (b *Blog) Update() (err error) {
	_, err = db.Exec("UPDATE blog SET title = ?, content = ? WHERE id = ?", b.Title, b.Content, b.Id)
	return
}

func (b *Blog) Query() (err error) {
	err = db.QueryRow("SELECT id, user_id, title, content, time FROM blog where id = ?", b.Id).Scan(&b.Id, &b.User_Id, &b.Title, &b.Content, &b.Time)
	return
}

func GetAllBlogs() (blogs []*Blog, err error) {
	rows, err := db.Query("SELECT id, user_id, title, content, time FROM blog")
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
