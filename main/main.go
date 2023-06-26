package main

import (
	"fmt"
	"murrou/main/api"
	"murrou/main/configer"
	"murrou/main/mapmanipulator"
)

func main() {

	// mapp := map[string]interface{}{
	// 	"url":     configer.GetConfigVal("example", "GET", "url", false),
	// 	"args":    configer.GetConfigVal("example", "GET", "args", true),
	// 	"headers": configer.GetConfigVal("example", "GET", "headers", true),
	// }
	// mapp := map[string]interface{}{
	// 	"url":     "https://api64.ipify.org",
	// 	"args":    map[string]string{"format": "json"},
	// 	"headers": map[string]string{"Content-Type": "application/json"},
	// }
	mapp := configer.GetSection("example", "GET")
	myIp, respCode := api.Get(mapp)
	// PARSING RESPONSE AND RESPONSE CODE
	fmt.Println("response code: ", respCode, "response body: ", myIp)

	// CONVERTING JSON STRING TO MAP
	myIpMap := mapmanipulator.JsonToMap(myIp)

	// DEMONSTRATING FOR LOOP
	for key, value := range myIpMap {
		fmt.Println("IP DATA: ", key, value)
	}

	// BUILDING STRUCT FOR POST REQUEST

	// type User struct {
	// 	Name  string
	// 	Email string
	// 	Id    string
	// }

	// // ASIGNING VALUES TO STRUCT
	// user := &User{Name: "Marin"}
	// user.Email = "marin@murrou.com"
	// user.Id = "1"

	// // CONVERTING STRUCT TO JSON
	// var jsonStruct = mapmanipulator.StructToJson(user)
	// // TEST THAT IS WORKING
	// fmt.Println("JSON FROM STRUCT: ", string(jsonStruct))

	// // MAKING POST REQUEST
	// postResponseBody, respCode := apicall.MakePostRequest(jsonStruct)

	// args := map[string]interface{}{
	// 	"url":      "https://httpbin.org",
	// 	"endpoint": "post",
	// 	"headers":  map[string]string{"Content-Type": "application/json"},
	// 	"body":     map[string]string{"name": "Marin", "email": "marin@abv.bg", "id": "1"},
	// }

	args := configer.GetSection("example", "POST")

	postResponseBody, respCode := api.Post(args)

	// SHOWING THE RESPONSE BODY AND RESPONSE
	fmt.Println("response code: ", respCode, "response body: ", postResponseBody)

	// CONVERTING REPONSE TO MAP
	postMap := mapmanipulator.JsonToMap(postResponseBody)

	// PROVING THE MAP

	for key, value := range postMap {
		fmt.Println("RESPONSE DATA :", key, value)
		if key == "data" {
			// LOWER LEVEL JSON VALUES
			fmt.Println("    ", "LOWER LEVEL FOR LOOP")
			// TYPE ASSERTION
			val := value.(string)
			v := mapmanipulator.JsonToMap(val)

			for key, value := range v {
				fmt.Println("    ", "    ", key, value)
			}
			fmt.Println("    ", "END OFLOWER LEVEL FOR LOOP")
			// END OF INTERNAL LOOP
		}
	}

}

// TEST
