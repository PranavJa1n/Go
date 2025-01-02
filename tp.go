package main

import (
	f "fmt"
	t "time"
)
func mainn(){
	var hello = 12
	f.Scan(&hello)
	f.Println(hello)
	f.Printf("%d\n",hello)
	f.Print(hello)
	f.Println(t.Now())
}
