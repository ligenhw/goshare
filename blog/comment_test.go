package blog

import (
	"log"
	"testing"
)

func TestComment(t *testing.T) {
	t.Log(CreateComment(58, 1, "good comment by gen"))
}

func TestCommentQueryByBlogId(t *testing.T) {
	log.SetFlags(log.Llongfile)
	comments, err := QueryByBlogId(10)
	if err != nil {
		t.Error(err)
	}
	for _, c := range comments {
		t.Log(*c)
	}

}
