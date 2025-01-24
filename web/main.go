package main

import (
	f "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func erCh(e error){
	if e!= nil{
		panic(e)
	}
}

func main(){
	f.Println("Hello")	
	var which string
	f.Println("Press G for get method and P for post method and PF for postform")
	f.Scanln(&which)
	if which == "G"{
		getResponse()
	} else if which == "P"{
		postResponse()
	} else if which == "PF" {
		postResponseForm()
	}
}

func getResponse(){
	const url = "http://127.0.0.1:1000/get"
	response, err := http.Get(url)
	erCh(err)
	defer response.Body.Close()
	var responseString strings.Builder
	responseByte, err := ioutil.ReadAll(response.Body)
	erCh(err)
	_ , err = responseString.Write(responseByte)
	erCh(err)
	f.Println(responseString.String())

}

func postResponse(){
	const url = "http://127.0.0.1:1000/post"
	// fake json payload
	jsonDataFake := strings.NewReader(`
		{
			"hello":"bye",
			"golang": "fun"
		}
	`)
	response, err := http.Post(url, "application/json", jsonDataFake)
	erCh(err)
	defer response.Body.Close()
	byteData , err := ioutil.ReadAll(response.Body)
	erCh(err)
	f.Println(string(byteData))
}

func postResponseForm(){
	const myurl = "http://127.0.0.1:1000/postform"

	data := url.Values{}

	data.Add("Firstname", "Pranav")
	data.Add("Lastname", "Jain")

	response, err := http.PostForm(myurl, data)
	erCh(err)
	content, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	f.Println(string(content))
}
