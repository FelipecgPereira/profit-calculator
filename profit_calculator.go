package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("Calculating profit...")

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)
	ratio := profit / revenue

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", expenses)
	fmt.Println("Taxa rate: ", ratio)

}
