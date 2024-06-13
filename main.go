package main

import (
	"Practice2024/may28"
	"Practice2024/sortPractice"
	"fmt"
)

func main() {
	average := may28.AvgSlice()
	fmt.Println("The average is:", average)

	avgEvenIndex := may28.EvenIndexedAverage()
	fmt.Println("The Avg of even Index number :  ", avgEvenIndex)

	avgOddIndex := may28.OddIndexedAverage()
	fmt.Println("The Avg of Odd Index number :  ", avgOddIndex)

	avgEvenSlice := may28.AvgEvenSlice()
	fmt.Println("The Avg of Even slice :  ", avgEvenSlice)

	avgOddSlice := may28.AvgOddSlice()
	fmt.Println("The Avg of Odd slice :  ", avgOddSlice)

	sortPractice.ArrSlice()
	sortPractice.SortArrWithOut()
	sortPractice.SortUsingFunction()

	flights := []sortPractice.Flight{
		{"Texas", 300},
		{"Arizona", 200},
		{"New York", 500},
		{"Vegas", 400},
	}

	fmt.Println("Before sorting:")
	sortPractice.PrintFlights(flights)

	sortedFlights := sortPractice.SortByPrice(flights)

	fmt.Println("\nAfter sorting by price (highest to lowest):")
	sortPractice.PrintFlights(sortedFlights)

}
