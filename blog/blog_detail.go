package blog

import (
	"github.com/ligenhw/goshare/user"
)

type BlogDetail struct {
	Blog `json:"blog"`
	User user.User `json:"user"`
}

func (b *BlogDetail) QueryByID() (err error) {
	err = db.QueryRow("select blog.id, user_id, title, content, blog.time,user.id, user_name, avatar_url from blog left join user on blog.user_id = user.id where blog.id = ?", b.Id).Scan(&b.Id, &b.UserId, &b.Title, &b.Content, &b.Time,
		&b.User.Id, &b.User.UserName, &b.User.AvatarUrl)
	return
}
