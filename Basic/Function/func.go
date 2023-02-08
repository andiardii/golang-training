package main

import "fmt"

func main() {
	/*
	 * anonymous function
	 */
	// Filtering := func(word string) bool {
	// 	return word == "Hot"
	// }

	FilterWord("Hot", Filtering)
	FilterWord("hot", Filtering)

	/*
	 * recursive function
	 */
	recur := FactorialRecursive(4)
	fmt.Println(recur)

}

/*
 * anonymous function
 */
func FilterWord(word string, filter Filter) {
	if filter(word) {
		fmt.Println("Teks Dilarang", word)
	} else {
		fmt.Println("Teks Diperbolehkan", word)
	}
}

type Filter func(word string) bool

func Filtering(word string) bool {
	return word == "Hot"
}

/*
 * recursive function
 */
 func FactorialRecursive(nilai int) int {
 	if nilai == 1 {
 		return 1
 	} else {
 		return nilai * FactorialRecursive(nilai-1)
 	}
 }