package config

import (
	"fmt"
	"reflect"
	"testing"
)

func printStruct(s interface{}, indent string) {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		if fieldVal.Kind() == reflect.Struct {
			fmt.Printf("%s%s:\n", indent, fieldType.Name)
			printStruct(fieldVal.Interface(), indent+"  ")
		} else {
			fmt.Printf("%s%s: %v\n", indent, fieldType.Name, fieldVal.Interface())
		}
	}
}

func TestParseLLpkgConfig(t *testing.T) {
	config, err := ParseLLpkgConfig("../../demo/llpkg.cfg")
	if err != nil {
		t.Errorf("Error parsing config file: %v", err)
	}
	printStruct(config, "")
}
