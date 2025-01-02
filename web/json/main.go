package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	About []string `json:"about,omitempty"`
}

func main() {
	fmt.Println("Hello")
	EncodeJson()
}

func EncodeJson(){
	userDetails := []user{
		{"Pranav", "xyz@gmail.com", "Hello", nil},
		{"Jain", "abc@gmail.com", "123", []string{"Linkedin", "golang is fun"}},
	}
	finalJson, err := json.MarshalIndent(userDetails, "", "\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",&finalJson)
}
