package simplenumber

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	list := make(map[int]int)

	for _, num := range nums {
		_, ok := list[num]
		if ok {
			list[num]++
			continue
		}
		list[num] = 1
	}

	for num, count := range list {
		if count == 1 {
			return num
		}
	}

	return 0
}

func singleNumberV2(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
