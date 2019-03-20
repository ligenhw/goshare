package user

import (
	"strconv"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	users, _ := GetAllUser()
	for _, u := range users {
		t.Log(*u)
	}
}

func TestCreateQueryDelete(t *testing.T) {
	u := User{UserName: "test1", Password: "testpass"}
	t.Log("Create result : ", u.Create())
	t.Log("Query result : ", u.Query())
	u.Password = "change123"
	t.Log("Update result : ", u.Update())
	t.Log(u)
	t.Log("Delete result: ", u.Delete())
}

func TestCreateExsts(t *testing.T) {
	u1 := User{UserName: "test2", Password: "testpass2"}
	u2 := u1
	t.Log("Create u : ", u1.Create())
	t.Log("Create u2 : ", u2.Create())
}

func TestI2S(t *testing.T) {
	s := string(37)
	t.Log(s)

	s1 := strconv.Itoa(37)
	t.Log(s1)
}
