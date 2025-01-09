package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

var age = 10

func main() {

	fmt.Println(age)
	//Declare a variable and use. Otherwise compiler error, to avoid thie error below is the way
	price := 10
	_ = price

	//Discover variable type
	startTime := time.Now()
	fmt.Printf("%T\n", startTime)

	//Get use input and print
	var name string
	fmt.Printf("Enter your name")
	fmt.Scanf("%s", &name)
	fmt.Println("Your name is: ", name)

	country := "australia:india"
	fmt.Println(utf8.RuneCountInString(country))
	fmt.Println(strings.Split(country, ":"))
}
