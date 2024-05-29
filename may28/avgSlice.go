package may28

var nums = []int{5, 7, 8, 2, 9, 4, 3, 2}

func AvgSlice() float64 {

	if len(nums) == 0 {
		return 0
	}
	sum = 0
	for _, num := range nums {
		sum += num
	}
	ret := float64(sum) / float64(len(nums))
	return ret

}

func AvgEvenSlice() float64 {
	if len(nums) == 0 {
		return 0
	}
	sum = 0
	count = 0
	for _, num := range nums {
		if num%2 == 0 {
			sum += num
			count++
		}
	}
	return float64(sum) / float64(count)
}

func AvgOddSlice() float64 {
	if len(nums) == 0 {
		return 0
	}
	sum = 0
	count = 0
	for _, num := range nums {
		if num%2 != 0 {
			sum += num
			count++
		}
	}
	return float64(sum) / float64(count)
}
