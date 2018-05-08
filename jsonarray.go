package gojson

import "bytes"

type Jsonarray struct {
	jva []JsonVar
}

func NewJsonArray() *Jsonarray {
	return &Jsonarray{}
}

func JsonArray(s string) (*Jsonarray, error) {
	ja := new(Jsonarray)
	j := new(jsonBase)
	err := j.Parse([]byte(s))
	if err != nil {
		return nil, err
	}
	m, err := j.Array()
	if err != nil {
		return nil, err
	}
	for _, val := range m {
		ja.jva = append(ja.jva, JsonVar{val})
	}
	return ja, nil
}

func (jo *Jsonarray) iniArrayTOList() []interface{} {
	data := jo.jva
	cnt := len(jo.jva)
	mm := make([]interface{}, cnt)
	for idx, ji := range data {
		mm[idx] = ji.data
	}
	return mm
}

func (ja *Jsonarray) GetValueByIndex(n int) *JsonVar {
	if n >= 0 && n < len(ja.jva) {
		return &ja.jva[n]
	} else {
		return nil
	}
}

func (ja *Jsonarray) Add(jv JsonVar) {
	ja.jva = append(ja.jva, jv)
}

func (ja *Jsonarray) Append(nja Jsonarray) {
	for _, val := range nja.jva {
		ja.Add(val)
	}
}

func (ja *Jsonarray) AddItem(val interface{}) {
	jv := &JsonVar{}
	if jo, ok := val.(*Jsonobject); ok {
		jv.data = jo.iniObjectTOMap()
	} else if njv, ok := val.(*JsonVar); ok {
		jv.data = njv.data
	} else if jo, ok := val.(*Jsonarray); ok {
		jv.data = jo.iniArrayTOList()
	} else {
		jv.data = val
	}
	ja.Add(*jv)
}

func (ja *Jsonarray) ToArray() []JsonVar {
	return ja.jva
}

func (ja *Jsonarray) setValue(inf []interface{}) {
	for _, val := range inf {
		jv := &JsonVar{val}
		ja.Add(*jv)
	}
}

func (jo *Jsonarray) GetJsonArrayItemList() []JsonVar {
	return jo.jva
}

func (ja *Jsonarray) Encode() string {
	b := bytes.Buffer{}
	b.WriteString("[")
	flag := false
	for _, key := range ja.jva {
		if flag {
			b.WriteString(",")
		}
		bo := bytes.Buffer{}
		val, err := key.JsonArray()
		if err != nil {
			if jov, err := key.Object(); err != nil {
				sov := key.AsString("")

				bo.WriteString(`"`)
				bo.WriteString(sov)
				bo.WriteString(`"`)
			} else {
				bo.WriteString(jov.Encode())
			}
		} else {
			if val == nil {
				bo.WriteString("")
			} else {
				bo.WriteString(val.Encode())
			}
		}
		flag = true
		b.WriteString(bo.String())
	}
	b.WriteString(`]`)
	return b.String()
}
