package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var i interface{}
	content1 := []byte(`{"id":"user1","name":"kenny","gender":"Male","age":23,"married":false,"cat": null,"dog": ["Amy", "Jack"]}`)
	err := json.Unmarshal(content1, &i)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
	m := i.(map[string]interface{})
	for k, v := range m {
		switch v.(type) {
		case float64:
			fmt.Printf("key=%s float value=%v\n", k, v)
		case bool:
			fmt.Printf("key=%s bool value=%v\n", k, v)
		case string:
			fmt.Printf("key=%s string value=%v\n", k, v)
		case nil:
			fmt.Printf("key=%s nil value=%v\n", k, v)
		case []interface{}:
			fmt.Printf("key=%s slice value=%v\n", k, v)
		}
	}
	// bool for JSON booleans,
	// float64 for JSON numbers,
	// string for JSON strings, and
	// nil for JSON null.

	var i2 interface{}
	content2 := []byte(`["a",1,false]`)
	err = json.Unmarshal(content2, &i2)
	if err != nil {
		panic(err)
	}
	fmt.Println(i2)
	// i2.([]interface{})
}
