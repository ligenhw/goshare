package blog

import "time"

type Comment struct {
	Id      string
	BlogId  string
	UserId  string
	Content string
	Time    time.Time
}

func (c *Comment) Create() (err error) {
	return nil
}
