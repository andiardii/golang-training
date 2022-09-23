package main

import "fmt"

func main(){
	nama := []string{
		"John", "Constantine", "Daredevils",
	}

	fmt.Println(nama)

	if nama[0] == "John" {
		fmt.Println("Ya ini adalah John")
	} else if nama[0] == "Constantine" {
		fmt.Println("Ini adalah " + nama[1])
	} else {
		fmt.Println("Hmm")
	}
}