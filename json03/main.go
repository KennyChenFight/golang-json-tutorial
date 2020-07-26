package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := map[string]string{
		"name":   "Kenny",
		"gender": "Male",
	}
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	var newUser map[string]string
	err = json.Unmarshal(b, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser)
}
