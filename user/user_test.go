package user

import "testing"

func TestAddUser(t *testing.T) {
	users := GetAllUser()
	t.Log("users ", users)
	id := AddUser(User{"user_test", "test", "123456", Profile{}})
	users = GetAllUser()
	t.Log("add user is ", users[id])
}
