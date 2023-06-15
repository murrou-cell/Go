package main

import (
	"fmt"
	"murrou/main/apicall"
	"murrou/main/mapmanipulator"
)

func main() {
	my_ip := apicall.Getmyip()

	my_ip_map := mapmanipulator.JsonToMap(my_ip)

	for key, value := range my_ip_map {
		fmt.Println(key, value)
	}

}
