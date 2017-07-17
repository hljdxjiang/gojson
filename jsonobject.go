package gojson

import (
	"bytes"
)

type Jsonobject struct {
	joi []jsonObjectItem
}

func NewJsonObject() *Jsonobject {
	return &Jsonobject{}
}

func JsonObject(s string) (*Jsonobject, error) {
	j := new(jsonBase)
	err := j.Parse([]byte(s))
	if err != nil {
		return nil, err
	}
	m, err := j.Map()
	if err != nil {
		return nil, err
	} else {
		return iniMaptoObject(m), nil
	}
}

func iniMaptoObject(m map[string]interface{}) *Jsonobject {
	js := new(Jsonobject)
	for key, val := range m {
		ja := &JsonVar{val}
		joi := &jsonObjectItem{key, *ja}
		js.joi = append(js.joi, *joi)
	}
	return js

}

func (jo *Jsonobject) iniObjectTOMap() map[string]interface{} {
	var mm map[string]interface{}
	mm = make(map[string]interface{})
	data := jo.joi
	for _, ji := range data {
		mm[ji.name] = ji.val.data
	}
	return mm
}

func (jo *Jsonobject) SetKey(name string, val interface{}) {
	ja := &JsonVar{}
	if jo, ok := val.(Jsonobject); ok {
		ja.data = jo.iniObjectTOMap()
	} else if jv, ok := val.(JsonVar); ok {
		ja.data = jv.data
	} else if jo, ok := val.(Jsonarray); ok {
		ja.data = jo.iniArrayTOList()
	} else {
		ja.data = val
	}
	br := jo.updateItem(name, *ja)
	if !br {
		nji := &jsonObjectItem{name, *ja}
		jo.joi = append(jo.joi, *nji)
	}
}

func (jo *Jsonobject) updateItem(name string, jv JsonVar) bool {
	for idx, ji := range jo.joi {
		if ji.name == name {
			jo.joi[idx].setVal(jv)
			return true
		}
	}
	return false
}

func (jo *Jsonobject) GetJsonItemList() []jsonObjectItem {
	return jo.joi
}

func (jo *Jsonobject) GetValue(n string) *JsonVar {
	ji := jo.lookup_item(n)
	if ji == nil {
		return nil
	}
	return &ji.val
}

func (jo *Jsonobject) GetValueByIndex(n int) *JsonVar {
	if n >= 0 && n < len(jo.joi) {
		return &jo.joi[n].val
	}
	return nil
}

func (jo *Jsonobject) Exists(n string) bool {
	if jo.lookup_item(n) != nil {
		return true
	}
	return false
}

func (jo *Jsonobject) lookup_item(n string) *jsonObjectItem {
	for _, ji := range jo.joi {
		if ji.name == n {
			return &ji
		}
	}
	return nil
}

func (jo *Jsonobject) AppendObject(njo Jsonobject) {
	for _, val := range njo.joi {
		ji := jo.lookup_item(val.name)
		if ji == nil {
			jo.joi = append(jo.joi, val)
		} else {
			ji.val = val.val
		}
	}
}

func (jo *Jsonobject) GetKeys() []string {
	var result []string
	for _, val := range jo.joi {
		result = append(result, val.name)
	}
	return result
}

func (jo *Jsonobject) GetIndex(name string) int {
	result := -1
	for idx, val := range jo.joi {
		if val.name == name {
			result = idx
			break
		}
	}
	return result
}

func (jo *Jsonobject) Encode() string {
	b := bytes.Buffer{}
	b.WriteString("{")
	flag := false
	for _, key := range jo.joi {
		if flag {
			b.WriteString(",")
		}
		bo := bytes.Buffer{}
		bo.WriteString(`"`)
		bo.WriteString(key.name)
		bo.WriteString(`":`)
		val, err := key.val.JsonArray()
		if err != nil {
			if jov, err := key.val.Object(); err != nil {
				sov := key.val.AsString("")
				switch key.val.data.(type) {
				case string, []byte:
					bo.WriteString(`"`)
					bo.WriteString(sov)
					bo.WriteString(`"`)
				default:
					bo.WriteString(sov)
				}

			} else {
				bo.WriteString(jov.Encode())
			}
		} else {
			bo.WriteString(val.Encode())
		}
		flag = true
		b.WriteString(bo.String())
	}
	b.WriteString("}")
	return b.String()
}
