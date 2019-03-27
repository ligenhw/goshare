package configration

import (
	"gopkg.in/yaml.v2"
)

func Parse(in []byte) (conf Configuration, err error) {
	err = yaml.Unmarshal(in, &conf)
	return
}
