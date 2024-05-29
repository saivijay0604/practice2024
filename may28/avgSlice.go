package may28

func AvgSlice() float64 {
	nums := []int{2, 6, 4, 2, 8}
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ret := float64(sum) / float64(len(nums))
	return ret

}
