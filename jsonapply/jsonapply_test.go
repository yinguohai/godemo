package jsonapply

import (
	"fmt"
	"testing"
)

func TestStructToJson(t *testing.T) {
	data := StructToJson()
	fmt.Println(string(data))
}

func TestMapToJson(t *testing.T) {
	data := MapToJson()
	fmt.Println(string(data))
}

func TestIntToJson(t *testing.T) {
	data := IntToJson()
	fmt.Println(string(data))
}

func TestJsonToStruct(t *testing.T) {
	JsonToStruct()
}

func TestJsonToMap(t *testing.T) {
	JsonToMap()
}
