package main

import (
	"murrou/main/apicall"
)

func main() {
	my_ip := apicall.Getmyip()
	println(my_ip)
}
