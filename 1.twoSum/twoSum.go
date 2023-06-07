package twosum

func twoSum(nums []int, target int) []int {

	// My first solution - Correction
	//  Time Complexity: O(n^2)
	//  Space Complexity: O(1)
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		sum := nums[i] + nums[j]
	// 		if sum == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// My second solution - Incorrect
	// for i := 0; i < len(nums)-1; i++ {
	// 	sum := nums[i] + nums[i+1]
	// 	if sum == target {
	// 		return []int{i, i + 1}
	// 	}
	// }

	/* solution correct and best performance
	Time Complexity: O(n)
	Space Complexity: O(n)
	*/
	idxMap := make(map[int]int)
	for idx, num := range nums {
		reqValue := target - num
		if reqIdx, ok := idxMap[reqValue]; ok {
			return []int{reqIdx, idx}
		}
		idxMap[num] = idx
	}

	return []int{}
}
