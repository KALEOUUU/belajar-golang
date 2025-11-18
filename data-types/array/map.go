package main

import "fmt"

func main() {
	// map with :=
	person := map[string]string{
		"name":    "kale morales",
		"address": "indonesia",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// map with make
	var books map[string]string = make(map[string]string) 

	books["title"] = "kepintaran orang goblok"
	books["author"] = "kale morales"
	books["year"] = "2025"

	fmt.Println(books) // mapping all key alphabet and value 
	fmt.Println(books["title"])
	fmt.Println(books["author"])

	delete(books, "year") // delete map key "year"
	fmt.Println(books)

}