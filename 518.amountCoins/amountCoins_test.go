package amountcoins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countChange(t *testing.T) {
	cases := map[string]struct {
		inputCoins   []int
		inputMoney   int
		expectedData int
	}{
		// "Expected 3 to money=4 and coins= [1,2,5]": {
		// 	inputCoins:   []int{1, 2, 5},
		// 	inputMoney:   5,
		// 	expectedData: 4,
		// },
		"": {
			inputCoins:   []int{3, 5, 7, 8, 9, 10, 11},
			inputMoney:   500,
			expectedData: 35502874,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			data := countChangeV2(cs.inputMoney, cs.inputCoins)

			assert.Equal(t, cs.expectedData, data)
		})
	}
}

func Benchmark_countChange(b *testing.B) {
	var (
		inputCoins = []int{3, 5, 7, 8, 9, 10, 11}
		inputMoney = 500
	)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countChange(inputMoney, inputCoins)
	}
}
func Benchmark_countChangeV2(b *testing.B) {
	var (
		inputCoins = []int{3, 5, 7, 8, 9, 10, 11}
		inputMoney = 500
	)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countChangeV2(inputMoney, inputCoins)
	}
}
