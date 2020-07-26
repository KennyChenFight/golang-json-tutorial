package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID string `json:"-"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Married bool `json:"married"`
}

type user2 struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
}

func main() {
	content1 := `{"id":"user1","name":"kenny","gender":"Male","age":23,"married":false`
	content2 := `{"id":"user1","name":"kenny","gender":"Male","age":23,"married":false}`
	fmt.Println(json.Valid([]byte(content1)))
	fmt.Println(json.Valid([]byte(content2)))

	var u1 user
	var u2 user2
	err := json.Unmarshal([]byte(content2), &u1)
	if err != nil {
		panic(err)
	}
	fmt.Println(u1)

	err = json.Unmarshal([]byte(content2), &u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(u2)
}