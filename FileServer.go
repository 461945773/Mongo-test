// FileServer.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("F:/Github/Mongo-test/"))
	err := http.ListenAndServe(":8000", fileServer)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}