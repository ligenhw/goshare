package blog

import "time"

type Blog struct {
	Id      string
	Titile  string
	Content string
	Time    time.Time
}

type Comment struct {
	Id      string
	BlogId  string
	UserId  string
	Content string
	Time    time.Time
}
