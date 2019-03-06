package user

import (
	"errors"
	"strconv"
	"time"
)

var UserList map[string]*User

func init() {
	UserList = make(map[string]*User)
	u := User{"user_1111", "gen", "123", Profile{"male", 28, "bj", "ligen@outlook.com"}}
	u2 := User{"user_2222", "root", "1234", Profile{"female", 22, "sh", "root@gmail.com"}}
	UserList["user_1111"] = &u
	UserList["user_2222"] = &u2
}

type User struct {
	Id       string
	UserName string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (*User, error) {
	if user, ok := UserList[uid]; ok {
		return user, nil
	}

	return nil, errors.New("User does not exists.")
}

func GetAllUser() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (*User, error) {
	if u, ok := UserList[uid]; ok {
		if uu.UserName != "" {
			u.UserName = uu.UserName
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}

	return nil, errors.New("User not exists")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.UserName == username && u.Password == password {
			return true
		}
	}

	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
