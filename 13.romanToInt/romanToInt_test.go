package romantoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToint(t *testing.T) {
	cases := map[string]struct {
		expectedData int
		input        string
	}{
		"should return 3": {
			expectedData: 3,
			input:        "III",
		},
		"should return 58": {
			expectedData: 58,
			input:        "LVIII",
		},
		"should return 1994": {
			expectedData: 1994,
			input:        "MCMXCIV",
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := romanToIntV3(cs.input)

			assert.Equal(t, cs.expectedData, resp)
		})
	}
}

func Benchmark_Patrick(b *testing.B) {
	numRoman := "MCMXCIV"
	for i := 0; i < b.N; i++ {
		romanToInt(numRoman)

	}
}
func Benchmark_PatrickV2(b *testing.B) {
	numRoman := "MCMXCIV"
	for i := 0; i < b.N; i++ {
		romanToIntV2(numRoman)

	}
}
func Benchmark_PatrickV3(b *testing.B) {
	numRoman := "MCMXCIV"
	for i := 0; i < b.N; i++ {
		romanToIntV3(numRoman)

	}
}
func Benchmark_PatrickV4(b *testing.B) {
	numRoman := "MCMXCIV"
	for i := 0; i < b.N; i++ {
		romanToIntV4(numRoman)

	}
}
func Benchmark_Google(b *testing.B) {
	numRoman := "MCMXCIV"
	for i := 0; i < b.N; i++ {
		GoogleV2(numRoman)
	}
}
