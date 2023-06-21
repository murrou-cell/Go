package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Get(args map[string]interface{}) (string, int) {

	// Building URL
	params := url.Values{}

	for k, v := range args {
		if strings.HasPrefix(k, "args") {
			ins := v.(map[string]string)
			for key, value := range ins {
				params.Add(key, value)
			}
		}
	}

	u, _ := url.ParseRequestURI(fmt.Sprint(args["url"]))
	if args["endpoint"] != nil {
		u.Path = fmt.Sprint(args["endpoint"])
	} else {
		u.Path = ""
	}
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	// Building request
	request, err := http.NewRequest("GET", urlStr, nil)
	checkErr(err)
	// Attaching headers
	for k, v := range args {
		if strings.HasPrefix(k, "headers") {
			ins := v.(map[string]string)
			for key, val := range ins {
				request.Header.Add(key, val)
			}
		}
	}

	// Executing request
	client := &http.Client{}
	response, err := client.Do(request)
	checkErr(err)

	responseBody, err := io.ReadAll(response.Body)
	checkErr(err)
	return string(responseBody), response.StatusCode
}

func Post(args map[string]interface{}) (string, int) {

	if args["body"] == nil {
		fmt.Println("WARNING: NO BODY ATTACHED TO POST REQUEST")
	}

	// Build URL
	u, _ := url.ParseRequestURI(fmt.Sprint(args["url"]))
	if args["endpoint"] != nil {
		u.Path = fmt.Sprint(args["endpoint"])
	} else {
		u.Path = ""
	}
	urlStr := fmt.Sprintf("%v", u)

	// Building body
	body, err := json.Marshal(args["body"])
	checkErr(err)
	// Building request
	request, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(body))
	checkErr(err)
	// Attaching headers
	for k, v := range args {
		if strings.HasPrefix(k, "headers") {
			ins := v.(map[string]string)
			for key, val := range ins {
				request.Header.Add(key, val)
			}
		}
	}
	// Executing request
	client := &http.Client{}
	response, err := client.Do(request)
	checkErr(err)

	responseBody, err := io.ReadAll(response.Body)
	checkErr(err)
	return string(responseBody), response.StatusCode
}
