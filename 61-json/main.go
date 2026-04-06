package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	marshal()
	unmarshal()
	encode()
	decode()
}

type User struct {
	Id          int      `json:"user_id"`
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`
	Permissions []string `json:"roles"`
}

func marshal() {
	u := User{
		Id:          1,
		Name:        "User One",
		Age:         20,
		Password:    "my-password",
		Permissions: []string{"admin", "group-member"},
	}

	var bytes []byte
	var err error
	bytes, err = json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshalling JSON")
		panic(err)
	}
	fmt.Println(string(bytes))
}

type Person struct {
	Name       string   `json:"full_name"`
	Age        int      `json:"years_old,omitempty"`
	Occupation string   `json:"-"`
	Languages  []string `json:"spoken_languages"`
}

func unmarshal() {
	jsonData := `{
		"full_name":"Jane Doe",
		"years_old":25,
		"spoken_languages": ["French", "German"]
	}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(err)
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Languages:", person.Languages)
	fmt.Println("Occupation:", person.Occupation)
}

type Person2 struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func encode() {
	p := Person2{
		Name:   "Alice",
		Age:    25,
		Active: true,
	}

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("error creating file:", err)
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(p)
	if err != nil {
		fmt.Println("error encoding person:", err)
		panic(err)
	}
}

func decode() {
	var person Person2
	file, err := os.Open("output.json")
	if err != nil {
		fmt.Println("error creating file:", err)
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("error decoding person:", err)
		panic(err)
	}

	fmt.Println(person)
}
