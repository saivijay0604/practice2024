package sortPractice

import (
	"fmt"
	"sort"
)

func ArrSlice() {
	a := [5]int{5, 6, 2, 3, 1}
	s := a[:]
	sort.Ints(s)
	fmt.Println("array sort", s)

}

func SortArrWithOut() {
	b := []string{"banana", "apple", "cherry", "date"}

	n := len(b)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}

		}
	}
	fmt.Println("sorted array with out function", b)

}
