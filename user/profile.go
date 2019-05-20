package user

import "github.com/ligenhw/goshare/orm"

type Profile struct {
	Id       int
	Openid   string
	AuthType int
	UserId   int
	Content  string
}

const (
	GOSHARE = iota
	GITHUB
	QQ
	ALIPAY
)

func init() {
	orm.RegisterModel(new(Profile))
}

func QueryUserIdByProfile(openid string) (userId int, err error) {
	p := &Profile{
		Openid: openid,
	}
	if err = o.Read(p, "openid"); err != nil {
		return
	}

	userId, err = p.UserId, nil
	return
}

func CreateProfile(authType, userid int, openid, content string) error {
	p := &Profile{
		Openid:   openid,
		AuthType: authType,
		UserId:   userid,
		Content:  content,
	}
	if _, err := o.Insert(p); err != nil {
		return err
	}

	return nil
}
