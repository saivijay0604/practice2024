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
}
