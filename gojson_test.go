package gojson

import (
	"fmt"
	"testing"

	"oldbaby.cn/gojson"
)

func JsonArrayTest(t *testing.T) {
	//new a json
	jo := NewJsonObject()

	//Json add value
	jo.SetKey("aa", "bb")
	jo.SetKey("bb", "cc")
	jo.SetKey("cc", 1)
	//Json Update value
	jo.SetKey("cc", true)

	//get string value with default value
	s := jo.GetValue("aa").AsString("sss")
	fmt.Println("GetValue ", s)
	//get string value
	ss, err := jo.GetValue("aa").String()
	fmt.Println(ss)
	if err != nil {
		fmt.Println("err is ", err)
	}
	//get bool value with default value
	bb := jo.GetValue("cc").AsBool(false)
	fmt.Println(bb)
	//get bool int64 with default value
	i := jo.GetValue("cc").AsInt(0)
	fmt.Println("print int ", i)
	fmt.Println("int to string", jo.GetValue("cc").AsString("cc is int"))
	fmt.Println("index 1", jo.GetValueByIndex(1).AsString("111"))

	ns := `{"a1":"a1","b1":"b1"}`
	//new json with string
	if njo, err := JsonObject(ns); err != nil {
		fmt.Println(err)
		//append anohter Json to this Json
		jo.AppendObject(*njo)
	} else {
		fmt.Println("JsonObject with string", njo.GetValue("a1").AsString())
	}

}

func JsonObjectTest(b *testing.B) {
	//new a jsonarray
	ja := NewJsonArray()
	//jsonarray add a new item
	ja.AddItem("aa")
	ja.AddItem("bb")

	jo := NewJsonObject()
	jo.SetKey("aa", "bb")
	njv := jo.GetValue("aa")
	//JsonArray add a jsonvar
	ja.Add(*njv)

	myjas := `["a","b"]`
	myja, err := gojson.JsonArray(myjas)
	//JsonArray append another JsonArray
	ja.Append(*myja)

	fmt.Println(ja.Encode())
	for idx, _ := range ja.ToArray() {
		fmt.Println("GetValueByIndex", ja.GetValueByIndex(idx).AsString())
	}
}

func ObjectWithArray(b *testing.B) {
	ss := `{"aa"：[{"aa":"bb"},{"bb":"cc"},{"cc":["a","b"]}]}`
	jo, err := JsonObject(ss)
	fmt.Println(err)
	ja, err := jo.GetValue("aa").JsonArray()
	njo, err := ja.GetValueByIndex(2).Object()
	nja, err := njo.GetValue("cc").JsonArray()
	fmt.Println(nja.GetValueByIndex(0).AsString())
	jas := `["t1","t2"]`
	myja, err := JsonArray(jas)
	fmt.Println(jo.Encode())
	jo.SetKey("ja", myja)
}
