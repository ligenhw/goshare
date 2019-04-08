package orm

import (
	"strings"
	"testing"
)

func TestSnakeString1(t *testing.T) {
	t1 := "UserName"
	t1r := snakeString(t1)
	if strings.Compare(t1r, "user_name") != 0 {
		t.Error(t1r)
	}
	t.Log(len("user_name"))
	t.Log(len(t1r))
}


func TestSnakeString(t *testing.T) {
	camel := []string{"PicUrl", "HelloWorld", "HelloWorld", "HelLOWord", "PicUrl1", "XyXX"}
	snake := []string{"pic_url", "hello_world", "hello_world", "hel_l_o_word", "pic_url1", "xy_x_x"}

	answer := make(map[string]string)
	for i, v := range camel {
		answer[v] = snake[i]
	}

	for _, v := range camel {
		res := snakeString(v)
		if res != answer[v] {
			t.Error("Unit Test Fail:", v, res, answer[v])
		}
	}
}

func TestStringEqual(t *testing.T) {
	t1 := "Use"
	t2 := string([]byte{'U', 's', 'e'})
	if t1 == t2 {
		t.Log("t1 == t2")
	} else {
		t.Log("t1 != t2")
	}
}
