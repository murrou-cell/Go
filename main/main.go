package main

import (
	"fmt"
	"murrou/main/apicall"
	"murrou/main/mapmanipulator"
)

func main() {
	// GET API CALL
	myIp, respCode := apicall.Getmyip()
	// PARSING RESPONSE AND RESPONSE CODE
	fmt.Println("response code: ", respCode, "response body: ", myIp)

	// CONVERTING JSON STRING TO MAP
	myIpMap := mapmanipulator.JsonToMap(myIp)

	// DEMONSTRATING FOR LOOP
	for key, value := range myIpMap {
		fmt.Println("IP DATA: ", key, value)
	}

	// BUILDING STRUCT FOR POST REQUEST

	type User struct {
		Name  string
		Email string
		Id    string
	}

	// ASIGNING VALUES TO STRUCT
	user := &User{Name: "Marin"}
	user.Email = "marin@murrou.com"
	user.Id = "1"

	// CONVERTING STRUCT TO JSON
	var jsonStruct = mapmanipulator.StructToJson(user)
	// TEST THAT IS WORKING
	fmt.Println("JSON FROM STRUCT: ", string(jsonStruct))

	// MAKING POST REQUEST
	postResponseBody, respCode := apicall.MakePostRequest(jsonStruct)

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
