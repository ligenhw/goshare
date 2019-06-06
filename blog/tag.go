package blog

import "github.com/ligenhw/goshare/orm"

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TagBlogRel struct {
	Id     int
	TagId  int
	BlogId int
}

func init() {
	orm.RegisterModel(new(Tag))
}

func CreateTag(name string) (id int64, err error) {
	t := &Tag{
		Name: name,
	}

	id, err = o.Insert(t)
	return
}

func deleteTag(id int) (err error) {
	t := &Tag{
		Id: id,
	}

	_, err = o.Delete(t)
	return
}

func GetTags() (tags []*Tag, err error) {
	tags = make([]*Tag, 0)
	qb := o.QueryTable(new(Tag))
	_, err = qb.OrderBy("-time").All(&tags)

	return
}

func AddTagsToBlog(blogId int, tagIds ...int) {
	db.Exec("")
}
