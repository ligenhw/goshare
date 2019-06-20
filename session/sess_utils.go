package session

import (
	"bytes"
	"encoding/gob"
)

// EncodeGob encode map data to gob
func EncodeGob(obj map[interface{}]interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(obj)
	if err != nil {
		return []byte(""), err
	}
	return buf.Bytes(), nil
}

// DecodeGob decode data to map
func DecodeGob(encoded []byte) (map[interface{}]interface{}, error) {
	buf := bytes.NewBuffer(encoded)
	decoder := gob.NewDecoder(buf)
	var out map[interface{}]interface{}
	err := decoder.Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
