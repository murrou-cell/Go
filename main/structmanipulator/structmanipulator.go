package structmanipulator

import (
	"fmt"
	"reflect"
)

func GenerateStruct(keys, types []string) interface{} {
	structType := reflect.StructOf(generateFields(keys, types))
	fmt.Println(reflect.New(structType).Elem().Interface())
	return reflect.New(structType).Elem().Interface()
}

func generateFields(keys, types []string) []reflect.StructField {
	fields := make([]reflect.StructField, len(keys))
	for i := 0; i < len(keys); i++ {
		fields[i] = reflect.StructField{
			Name: keys[i],
			Type: getType(types[i]),
		}
	}
	return fields
}

func getType(typeStr string) reflect.Type {
	switch typeStr {
	case "string":
		return reflect.TypeOf("")
	case "int":
		return reflect.TypeOf(0)
	default:
		panic(fmt.Sprintf("Unknown type: %s", typeStr))
	}
}
