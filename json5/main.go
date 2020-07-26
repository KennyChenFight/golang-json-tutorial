package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID string `json:"-"`
	Name string `json:"name"`
	Gender string `json:"gender,omitempty"`
	Age int `json:"-,"`
	Married bool `json:",omitempty"`
}

func main() {
	originalUser := user{
		ID:       "user1",
		Name:     "kenny",
		Gender: "",
		Age: 23,
		Married: true,
	}
	b, err := json.Marshal(originalUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
}

