package may28

func EvenIndexedAverage() float64 {
	nums := []int{5, 2, 4, 2, 7, 3}
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	count := 0
	for i, num := range nums {
		if i%2 == 0 {
			sum += num
			count++
		}
	}
	ret := float64(sum) / float64(count)

	return ret
}
