package blog

import "github.com/ligenhw/goshare/orm"

type Tag struct {
	Id   int
	Name string
}

type BlogTagRel struct {
	Id     int
	BlogId int
	TagId  int
}

func init() {
	orm.RegisterModel(new(Tag))
}

func CreateTag(name string) (err error) {
	t := &Tag{
		Name: name,
	}

	_, err = o.Insert(t)
	return
}
