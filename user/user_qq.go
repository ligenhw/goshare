package user

import (
	"encoding/json"
	"log"

	"github.com/ligenhw/goshare/orm"
)

type QqUser struct {
	Openid  string `orm:"pk"`
	UserId  int
	Content string
}

func init() {
	orm.RegisterModel(new(QqUser))
}

func QueryUserByQQ(openid string) (user *User, err error) {
	qquser := QqUser{
		Openid: openid,
	}
	if err = o.Read(&qquser); err != nil {
		return
	} else {
		log.Println("have this openid, query :", openid)
		user = &User{
			Id: qquser.UserId,
		}
		if err = user.QueryByID(); err != nil {
			log.Println(err)
			return
		}
	}
	return
}

func CreateUserByQQ(openid, content string) (user *User, err error) {
	var qqinfo struct {
		NickName  string `json:"nickname"`
		AvatarUrl string `json:"figureurl_qq"`
	}
	if err = json.Unmarshal([]byte(content), &qqinfo); err != nil {
		return
	}

	user = &User{
		UserName:  qqinfo.NickName,
		AvatarUrl: qqinfo.AvatarUrl,
	}
	if err = user.Create(); err != nil {
		return
	}

	log.Println("user create success id : ", user.Id)

	qquser := QqUser{
		Openid:  openid,
		UserId:  user.Id,
		Content: content,
	}

	if _, err = o.Insert(qquser); err != nil {
		return
	}
	log.Println("qq user create success")

	return
}
