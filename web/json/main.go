package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	About    []string `json:"about,omitempty"`
}

func main() {
	// fmt.Println("Hello")
	// EncodeJson()
	ConsumingJson()
}

func EncodeJson() {
	userDetails := []user{
		{"Pranav", "xyz@gmail.com", "Hello", nil},
		{"Jain", "abc@gmail.com", "123", []string{"Linkedin", "golang is fun"}},
	}
	finalJson, err := json.MarshalIndent(userDetails, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", &finalJson)
}

func ConsumingJson() {
	jsonData := []byte(`
	{
	"username": "Pranav",
	"email": "xyz@gmail.com",
	"about": ["Linkedin","golang is fun"]
	}
	`)

	// using struct

	var myUserData user
	checkValid := json.Valid(jsonData)
	if checkValid {
		json.Unmarshal(jsonData, &myUserData)
		fmt.Printf("%#v\n\n", myUserData)
	} else {
		fmt.Println("JSON data not valid")
	}

	// If I want data in key value pair

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	for k, v := range myData {
		fmt.Printf("Key: %v\tValue: %v\n", k, v)
	}
}
