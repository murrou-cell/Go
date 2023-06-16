package mapmanipulator

import (
	"encoding/json"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func JsonToMap(jsonS string) map[string]interface{} {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonS), &data)
	checkErr(err)

	return data

}

func StructToJson(s any) []byte {
	j, err := json.Marshal(s)
	checkErr(err)
	return j
}
