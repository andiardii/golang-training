package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "Rossinante"
	nama[1] = "Constantine"
	nama[2] = "Donquixote"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var nilai = [3]int{
		80,
		85,
		90,
	}

	fmt.Println(nilai)
	fmt.Println(nilai[0])
	fmt.Println(len(nilai))
}