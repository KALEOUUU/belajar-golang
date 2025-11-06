// membuat program sederhana untuk menampilkan data diri dengan variable dan constant

package main

import "fmt"

func main() {

	type School string

	const (
		firstName = "kale"
		lastName = "morales"
	)

	age := 17

	latestSchool := School("High School") // konversi tipe data string ke tipe data kustom School

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Age:", age)
	fmt.Println("Latest School:", latestSchool)
}