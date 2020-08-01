package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type user struct {
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *user) MarshalJSON() ([]byte, error) {
	//createdAt := u.CreatedAt.Format("2006-01-02 15:04:05")
	//return []byte(fmt.Sprintf(`{"name":"%s","gender":"%s","createdAt":"%s"}`, u.Name, u.Gender, createdAt)), nil

	type AliasUser user
	return json.Marshal(&struct {
		CreatedAt string `json:"createdAt"`
		*AliasUser
	}{
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		AliasUser: (*AliasUser)(u),
	})
}

func (u *user) UnmarshalJSON(data []byte) error {
	type AliasUser user
	au := struct {
		CreatedAt string `json:"createdAt"`
		*AliasUser
	}{
		AliasUser: (*AliasUser)(u),
	}

	if err := json.Unmarshal(data, &au); err != nil {
		return err
	}
	newCreatedAt, err := time.ParseInLocation("2006-01-02 15:04:05", au.CreatedAt, time.Local)
	if err != nil {
		return err
	}

	u.CreatedAt = newCreatedAt
	return nil
}

func main() {
	u := user{
		Name:      "Kenny",
		Gender:    "Male",
		CreatedAt: time.Now(),
	}
	b, err := json.Marshal(&u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	var u2 user
	err = json.Unmarshal(b, &u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(u2)
}
