package main

import "fmt"

func main() {
	var name string // deklarasi variable dengan tipe data string

	name = "kale morales" // assignment nilai ke variable
	fmt.Println(name)

	name = "kaleee"
	fmt.Println(name)

	var (
		firstName = "kale"
		lastName  = "morales"
	) // deklarasi multiple variable

	fmt.Println(firstName)
	fmt.Println(lastName)

	var age int32 // deklarasi variable dengan tipe data int32
	
	age = 20
	fmt.Println(age)

	var (
		height = 180
		weight = 60
	)

	fmt.Println(height)
	fmt.Println(weight)

	friendName := "joko" // variable dengan ":="
	fmt.Println(friendName)

	friendAge := 25
	fmt.Println(friendAge)

	country := "indonesia" // variable dengan ":="
	fmt.Println(country)

	country = "United States" // variable yang sudah dideklarasi tidak bisa dideklarasi ulang dengan ":="
	fmt.Println(country)

	// country := "Amerika Serikat" // ini akan error
	// fmt.Println(country)

}