// HEAD.go

package main

import (
	"fmt"
	"net/http"w
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage :", os.Args[0], "host:port")
		os.Exit(1)
	}
	url := os.Args[1]
	url = "http://" + url + ":80	"
	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}
	os.Exit(0)
}
