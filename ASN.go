// SaveJson

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name:  Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"}, Email{Kind: "work", Address: "j.newmarch@boxhill.edu.cn"}},
	}
	saveJSON("person.json", person)
}

func saveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	defer outFile.Close()
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}
