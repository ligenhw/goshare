package blog

import "testing"

func TestComment(t *testing.T) {
	t.Log(Create(58, 1, "good comment by gen"))
}

func TestReadComment(t *testing.T) {

}

func TestCommentQueryByBlogId(t *testing.T) {
	comments, err := QueryByBlogId(55)
	if err != nil {
		t.Error(err)
	}
	for _, c := range comments {
		t.Log(*c)
	}
	
}
