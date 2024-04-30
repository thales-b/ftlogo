package main

import "fmt"

// Customer represents information about a customer.
type Customer struct {
	Name  string
	Email string
}

func main() {
	var title string
	var copies int
	var author string
	var numEdition int
	var isDiscount bool
	var discountPercent float64
	title = "For the Love of Go"
	author = "Jean de Meung"
	copies = 99
	numEdition = 3
	isDiscount = true
	discountPercent = .2
	fmt.Println(title)
	fmt.Println(author)
	fmt.Println(copies)
	fmt.Println(numEdition)
	fmt.Println(isDiscount)
	fmt.Println(discountPercent)
}
