package gojson

import (
	"encoding/json"
	"errors"
	"strconv"
)

type JsonVar struct {
	data interface{}
}

func (jv *JsonVar) Value() interface{} {
	return jv.data
}

func (jv *JsonVar) String() (string, error) {
	if jv == nil {
		return "", errors.New("this is not a jsonvar")
	}
	s := ""
	data := jv.data
	switch jv.data.(type) {
	case bool:
		s = strconv.FormatBool(data.(bool))
		break
	case float32:
		s = strconv.FormatFloat(float64(data.(float32)), 'f', -1, 64)
		break
	case float64:
		s = strconv.FormatFloat(data.(float64), 'f', -1, 64)
		break
	case int:
		s = strconv.FormatInt(int64(data.(int)), 10)
		break
	case int16:
		s = strconv.FormatInt(int64(data.(int16)), 10)
		break
	case int32:
		s = strconv.FormatInt(int64(data.(int32)), 10)
		break
	case int64:
		s = strconv.FormatInt(data.(int64), 10)
		break
	case int8:
		s = strconv.FormatInt(int64(data.(int8)), 10)
		break
	case uint:
		s = strconv.FormatUint(uint64(data.(uint)), 10)
		break
	case uint16:
		s = strconv.FormatUint(uint64(data.(uint16)), 10)
		break
	case uint32:
		s = strconv.FormatUint(uint64(data.(uint32)), 10)
		break
	case uint64:
		s = strconv.FormatUint(data.(uint64), 10)
		break
	case uint8:
		s = strconv.FormatUint(uint64(data.(uint8)), 10)
		break
	case string:
		s = data.(string)
		break
	case []byte:
		s = string(data.([]byte))
		break
	case json.Number:
		x, _ := data.(json.Number)
		s = string(x)
		break
	default:
		return "", errors.New("type assertion to string faild")
	}
	return s, nil
}

func (jv *JsonVar) Bool() (bool, error) {
	if jv == nil {
		return false, errors.New("this is not a jsonvar")
	}
	if val, err := jv.String(); err == nil {
		if val == "true" {
			return true, nil
		}
	}
	return false, errors.New("type assertion to bool faild")

}

func (jv *JsonVar) Byte() ([]byte, error) {
	if jv == nil {
		return nil, errors.New("this is not a jsonvar")
	}
	if s, err := jv.String(); err == nil {
		return []byte(s), nil
	} else {
		return nil, err
	}
}

func (jv *JsonVar) Array() ([]interface{}, error) {
	if jv == nil {
		return nil, errors.New("this is not a jsonvar")
	}
	if s, ok := jv.data.([]interface{}); ok {
		return s, nil
	}
	return nil, errors.New("type assertion to Array faild")
}

func (jv *JsonVar) JsonArray() (*Jsonarray, error) {
	if jv == nil {
		return nil, errors.New("this is not a jsonvar")
	}
	ja := &Jsonarray{}
	arr, err := jv.Array()
	if err == nil {
		ja.setValue(arr)
		return ja, nil
	}
	return nil, errors.New("type assertion to Array faild")
}

func (jv *JsonVar) Object() (*Jsonobject, error) {
	if jv == nil {
		return nil, errors.New("this is not a jsonvar")
	}
	s, err := jv.Map()
	if err == nil {
		return iniMaptoObject(s), nil
	}
	return nil, errors.New("type assertion to Object faild")
}

func (jv *JsonVar) Map() (map[string]interface{}, error) {
	if jv == nil {
		return nil, errors.New("this is not a jsonvar")
	}
	if s, ok := (jv.data).(map[string]interface{}); ok {
		return s, nil
	}
	return nil, errors.New("type assertion to Map faild")
}

func (jv *JsonVar) Int64() (int64, error) {
	if jv == nil {
		return 0, errors.New("this is not a jsonvar")
	}
	if val, err := jv.String(); err == nil {
		if nval, nerr := strconv.ParseInt(val, 10, 64); nerr == nil {
			return nval, nil
		}
	}
	if s, ok := jv.data.(int64); ok {
		return s, nil
	}
	return 0, errors.New("type assertion to int64 faild")
}

func (jv *JsonVar) Float64() (float64, error) {
	if jv == nil {
		return 0.0, errors.New("this is not a jsonvar")
	}
	if val, err := jv.String(); err == nil {
		if nval, nerr := strconv.ParseFloat(val, 64); nerr == nil {
			return nval, nil
		}
	}
	return 0.0, errors.New("type assertion to float64 faild")
}

func (jv *JsonVar) AsString(args ...string) string {
	def := ""
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
		break
	}
	r, err := jv.String()
	if err == nil {
		return r
	}
	return def
}

func (jv *JsonVar) AsInt(args ...int) int64 {
	def := 0
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
		break
	}
	if i, err := jv.Int64(); err == nil {
		return i
	}
	return int64(def)
}

func (jv *JsonVar) AsBool(args ...bool) bool {
	def := false
	switch len(args) {
	case 0:
		break
	case 1:
		def = args[0]
		break
	}
	if i, err := jv.Bool(); err != nil {
		return i
	}
	return def
}
