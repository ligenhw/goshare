package blog

import "testing"

func TestComment(t *testing.T) {
	t.Log(Create(58, 1, "good comment by gen"))
}
