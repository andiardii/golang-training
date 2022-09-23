package main

import "fmt"

func main() {
	angka := []int{
		1, 2, 3, 4,
	}

	fmt.Println(angka)

	switch angka[0]{
	case 1:
		fmt.Println(angka[0])
	case 2:
		fmt.Println(angka[1])
	default:
		fmt.Println(angka)
	}

	//support short statement
	//support switch empty
	switch {
	case angka[1] == 1:
		fmt.Println(angka[0])
	case angka[0] == 2:
		fmt.Println(angka[1])
	default:
		fmt.Println(angka)
	}
}