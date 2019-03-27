package configration

import (
	"io/ioutil"
	"testing"
)

func TestConfiguration(t *testing.T) {
	file := "config.yaml"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	config, err := Parse(content)
	if err != nil {
		panic(err)
	}

	t.Log(config)
}
