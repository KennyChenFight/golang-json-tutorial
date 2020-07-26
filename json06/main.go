package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
	Cat     *cat
}

type cat struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func main() {
	originalUser := user{
		ID:      "user1",
		Name:    "kenny",
		Gender:  "Male",
		Age:     23,
		Married: false,
		Cat:     nil,
	}
	b, err := json.Marshal(originalUser)
	//b, err := json.MarshalIndent(originalUser, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
}
