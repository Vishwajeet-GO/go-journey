package main

import "fmt"

func main() {

	var maths int
	fmt.Println("Enter the marks of Maths out of 100:")
	fmt.Scanln(&maths)

	var science int
	fmt.Println("Enter the marks of Science out of 100:")
	fmt.Scanln(&science)

	var english int
	fmt.Println("Enter the marks of English out of 100:")
	fmt.Scanln(&english)

	var history int
	fmt.Println("Enter the marks of History out of 100:")
	fmt.Scanln(&history)

	var hindi int
	fmt.Println("Enter the marks of Hindi out of 100:")
	fmt.Scanln(&hindi)

	total := maths + science + english + history + hindi

	fmt.Printf("The total marks obtained is: %d out of 500\n", total)

	percentage := (float64(total) / 500.0) * 100
	fmt.Printf("The percentage is: %.2f%%\n ", percentage)

	if percentage >= 90 {
		fmt.Println("Grade: A")
	} else if percentage >= 80 {
		fmt.Println("Grade: B")
	} else if percentage >= 70 {
		fmt.Println("Grade: C")
	} else if percentage >= 60 {
		fmt.Println("Grade: D")
	} else {
		context := "Better luck next time!"
		fmt.Println("Grade F:", context)
	}
}
