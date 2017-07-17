package gojson

import (
	"bytes"
	"encoding/json"
	"errors"
)

type jsonBase struct {
	data interface{}
}

func (jb *jsonBase) Parse(b []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(b))
	dec.UseNumber()
	return dec.Decode(&jb.data)
}

func (jb *jsonBase) Map() (map[string]interface{}, error) {
	if m, ok := jb.data.(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map failed")

}

func (jb *jsonBase) Array() ([]interface{}, error) {
	if m, ok := jb.data.([]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map failed")
}

func (jb *jsonBase) Encode() (string, error) {
	b, err := json.Marshal(&jb.data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
