package searchinsertposition

func searchInsert(nums []int, target int) int {
	for idx, value := range nums {
		if target <= value {
			return idx
		}
	}
	return len(nums)
}

func searchInsertV2(nums []int, target int) int {
	start, end, mid := 0, len(nums)-1, 0

	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
