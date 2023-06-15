package apicall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")
	checkErr(err)
	d := out.Bytes()
	return string(d)
}

func readBody(resp http.Response) string {
	responseBody, err := io.ReadAll(resp.Body)
	checkErr(err)
	return formatJSON(responseBody)
}

func Getmyip() string {
	request, err := http.NewRequest("GET", "https://api64.ipify.org?format=json", nil)
	checkErr(err)

	client := &http.Client{}
	response, err := client.Do(request)
	checkErr(err)

	return readBody(*response)
}
