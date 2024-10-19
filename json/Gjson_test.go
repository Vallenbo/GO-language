package json

import (
	"github.com/tidwall/gjson"
	"testing"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

type name struct {
}

func Test_gjson(*testing.T) {

	value := gjson.Get(json, "name.last")
	println(value.String())
}
