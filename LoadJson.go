// LoadJson.go
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
	var person Person
	loadJSON("person.json", &person)
	fmt.Println("Person :\n", person)
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ":" + v.Address
	}
	return s
}

func loadJSON(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	defer inFile.Close()
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}
