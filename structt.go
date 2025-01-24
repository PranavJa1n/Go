package main

import f "fmt"

type stock struct{
	Chocolate int
	Coffee int
}

func main(){
	f.Println("Hello");
	past := stock{2000, 70};
	f.Println("Total Chocolate in past are",past.Chocolate);
	f.Println("Total Coffee in past",past.Coffee);
	var p int;
	f.Scan(&p);
	f.Println(past.price(p)) 
}

func (test stock) price(price int) int{
	var total int = test.Chocolate * price;
	return total;
}
