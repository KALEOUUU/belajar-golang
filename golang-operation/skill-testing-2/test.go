// exercise golang operation skill-testing-2

package main

import "fmt"

func main() {
	math := 40
	english := 70
	cs := 90

	// lets total and find average

	total := math + english + cs
	average := total / 3

	fmt.Println("Total:", total)
	fmt.Println("Average:", average)

	passAverage := 60

	fmt.Println("Lulus:", average >= passAverage)
}