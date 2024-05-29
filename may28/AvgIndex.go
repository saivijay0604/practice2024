package may28

var (
	sum   = 0
	count = 0
)

func EvenIndexedAverage() float64 {
	if len(nums) == 0 {
		return 0
	}
	sum = 0
	count = 0
	for i, num := range nums {
		if i%2 == 0 {
			sum += num
			count++
		}
	}
	ret := float64(sum) / float64(count)

	return ret
}

func OddIndexedAverage() float64 {
	if len(nums) == 0 {
		return 0
	}
	sum = 0
	count = 0
	for i, num := range nums {
		if i%2 != 0 {
			sum += num
			count++
		}
	}
	return float64(sum) / float64(count)
}
