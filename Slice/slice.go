package main

import "fmt"

func main() {
	var nama = [...]string{
		"Constantine",
		"Roscinante",
		"Percival",
		"Trafalgar",
		"Sherlock",
		"Edogawa",
	}

	fmt.Println(nama[2:4])
	fmt.Println(cap(nama))

	nama[2] = "Idra"
	fmt.Println(nama)

	//array last
	var tambahExt = append(nama[5:], "Conan")
	fmt.Println(tambahExt)
	fmt.Println(nama)

	tambahExt[0] = "Shinichi"
	fmt.Println(tambahExt)
	fmt.Println(nama)

	//array middle
	var tambahInt = append(nama[2:4], "Thor")
	fmt.Println(tambahInt)
	fmt.Println(nama)

	tambahInt[1] = "Daenerys"
	fmt.Println(tambahInt)
	fmt.Println(nama)

	//Syntax => make(typeData, length, capacity)
	data := make([]string, 4, 4)
	data[0] = "Viserys"
	data[1] = "Targaryen"
	data[2] = "Ragnarok"
	data[3] = "Zeus"

	fmt.Println(data)

	//copy(destination, source)
	salin := make([]string, len(data), cap(data))
	copy(salin, data)
	fmt.Println(salin)

	// array [...]
	// slice []
}