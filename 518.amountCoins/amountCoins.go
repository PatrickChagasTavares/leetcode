package amountcoins

func countChange(money int, coins []int) int {
	if money == 0 {
		return 1
	}

	if money < 0 || len(coins) == 0 {
		return 0
	}

	return countChange(money-coins[0], coins) + countChange(money, coins[1:])
}

func countChangeV2(money int, coins []int) int {
	// for each amount, there is sum {amount - coins[i]} methods
	result := make([]int, money+1)
	// base case
	result[0] = 1
	for i := 0; i < len(coins); i++ {
		// think about to use the ith-coin
		// here don't need to consider about duplicated combinations
		// because ith coin is introduced and the rest does not
		for j := 1; j <= money; j++ {
			// add on the method
			if j >= coins[i] {
				result[j] += result[j-coins[i]]
			}
		}
	}
	return result[money]
}
