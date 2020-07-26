package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const jsonStream = `
	[
		{"name":"kenny","gender":"male"},
		{"name": "nicole","gender": "female"}
	]`
	type user struct {
		Name string
		Gender string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	t, err := dec.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf("type: %T, t: %v\n", t, t)

	for dec.More() {
		var u user
		err := dec.Decode(&u)
		if err != nil {
			panic(err)
		}
		fmt.Println(u)
	}

	t, err = dec.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf("type: %T, t: %v\n", t, t)
}
