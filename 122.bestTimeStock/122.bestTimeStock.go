package main

import "fmt"

func maxProfit(prices []int) int {
	var s = prices[len(prices)-1]
	var maxP = 0

	for i := len(prices) - 2; i > 0; i-- {
		if prices[i] < s {
			maxP += s - prices[i]
			s = prices[i]
			continue
		}
		s = prices[i]
	}
	return maxP
}

func main() {
	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
