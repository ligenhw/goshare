package blog

import "testing"

func TestTag(t *testing.T) {
	tagNames := [...]string{"golang", "源码"}
	tagIds := [len(tagNames)]int64{}
	var err error
	if tagIds[0], err = CreateTag(tagNames[0]); err != nil {
		t.Error(err)
	}
	if tagIds[1], err = CreateTag(tagNames[1]); err != nil {
		t.Error(err)
	}

	if tags, err := GetTags(); err != nil {
		t.Error(err)
	} else {
		allTagsName := make([]string, 0)
		for _, tag := range tags {
			allTagsName = append(allTagsName, tag.Name)
		}

		for _, tag := range tagNames {
			if !contain(allTagsName, tag) {
				t.Error("get tags not contain : " + tag)
			}
		}
	}

	for _, id := range tagIds {
		if err = deleteTag(int(id)); err != nil {
			t.Error(err)
		}
	}
}

func contain(lists []string, item string) bool {
	for _, l := range lists {
		if l == item {
			return true
		}
	}

	return false
}
