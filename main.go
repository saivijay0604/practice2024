package main

import (
	"Practice2024/may28"
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

}
