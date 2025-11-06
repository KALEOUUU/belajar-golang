package main

import "fmt"

func main() {
	const name = "kale morales" // const untuk nilai yang tetap dan tidak bisa diubah
	const age = 25 
	const values = 100.5

	// fmt.Println("name =", name) // const tidak akan error jika tidak digunakan
	fmt.Println("age =", age)
	fmt.Println("values =", values)

	// multiple const
	const (
		firstName = "kale"
		lastName  = "morales"
	)

	fmt.Println("firstName =", firstName)
	fmt.Println("lastName =", lastName)
}