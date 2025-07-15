package main

import "fmt"

func main() {
	var name string
	fmt.Println("Howdy! What's your name?")
	fmt.Scan(&name)
	fmt.Printf("Hi, %s! Nice to meet you!\n", name)
}
