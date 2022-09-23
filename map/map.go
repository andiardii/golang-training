package main

import "fmt"

func main() {
	//map[key_type]value_type
	customer := map[string]string{
		"Names" : "Constantine",
		"Address" : "House of Mystery",
	}

	customer["partner"] = "Zatanna"
	customer["enemy"] = "Ghost"

	fmt.Println(customer)
	fmt.Println(customer["Names"] + " lives in " + customer["Address"] + ", fall in love with " + customer["partner"])

	//delete(variable, "key_index")
	delete(customer, "enemy")
	fmt.Println(customer)
}