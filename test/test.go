package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Id        string   `json:"id"`
	Name      []string `json:"name"`
	UserId    string   `json:"user_id"`
	CreatedAt int64    `json:"created_at"`
}

func main() {
	name := []string{"you"}
	m := make(map[string]interface{})
	m["id"] = "2"
	m["name"] = name
	m["user_id"] = "123"
	m["created_at"] = 5
	// v := make(map[string]string)
	// v = map[Address:[Mbezi juu,jogoo] Category:Staff City:[Dar es salaak] Code:[14128] Contact: Country:[Tanzania] Email:peterkelvin16@gmail.com FirstName:[] Image: LastName:[Mtera] Password:admin Title:admin@dot.com id:peterkelvin16@gmail.com]

	jsonString, _ := json.Marshal(m)
	fmt.Println(string(jsonString))

	// convert json to struct
	s := MyStruct{}
	json.Unmarshal(jsonString, &s)
	fmt.Println(s)
}
