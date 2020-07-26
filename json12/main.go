package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type CustomDateTimeFormat string

func (c CustomDateTimeFormat) rfc3339Format() string {
	return time.RFC3339
}

func (c CustomDateTimeFormat) myFormat() string {
	return "2006-01-02 15:04:05"
}

type CustomDateTime struct {
	time.Time
	CustomDateTimeFormat
}

func (d *CustomDateTime) MarshalJSON() ([]byte, error) {
	dateTimeStr := d.Time.Format(d.CustomDateTimeFormat.myFormat())
	return []byte(fmt.Sprintf(`"%s"`, dateTimeStr)), nil
}

func (d *CustomDateTime) UnmarshalJSON(b []byte) error {
	b = bytes.Trim(b, `"`)
	dataTime, err := time.Parse(d.CustomDateTimeFormat.myFormat(), string(b))
	if err != nil {
		return err
	}
	d.Time = dataTime
	return nil
}

type User struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	CreatedAt CustomDateTime `json:"createdAt"`
}

func main() {
	u := User{
		Name:      "Kenny",
		Gender:    "Male",
		CreatedAt: CustomDateTime{
			Time:   time.Now(),
		},
	}
	b, err := json.Marshal(&u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	var u2 User
	err = json.Unmarshal(b, &u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(u2)
}
