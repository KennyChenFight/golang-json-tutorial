package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	content := []string{"a", "b", "c"}
	b, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	var newContent []string
	err = json.Unmarshal(b, &newContent)
	if err != nil {
		panic(err)
	}
	fmt.Println(newContent)
}
