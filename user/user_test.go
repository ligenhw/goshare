package user

import (
	"testing"
)

func TestUser(t *testing.T) {
	u, err := CreateTestUser(t)
	if err != nil {
		t.Error("can not create a test user")
	}

	u2 := User{Id: u.Id}
	err = u2.QueryByID()
	if err != nil {
		t.Error(err)
	}

	if u.UserName != u2.UserName || u.Password != u2.Password {
		t.Error("Query User Error")
	}

	users, _ := GetAllUser()

	err = u2.Delete()
	if err != nil {
		t.Error("Delete User failed.")
	}

	users2, _ := GetAllUser()
	if len(users)-len(users2) != 1 {
		t.Error("After delete count do not decrease")
	}
}

func CreateTestUser(t *testing.T) (user *User, err error) {
	u := User{UserName: "testuser", Password: "testpass"}
	// clean up
	u.QueryByName()
	u.Delete()

	err = u.Create()
	if err != nil {
		t.Error(err)
		return
	}

	user = &u
	t.Log("user create success")
	return
}
