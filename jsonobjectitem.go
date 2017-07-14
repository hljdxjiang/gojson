package gojson

type jsonObjectItem struct {
	name string
	val  JsonVar
}

func (ji *jsonObjectItem) setVal(jv JsonVar) {
	ji.val = jv

}
