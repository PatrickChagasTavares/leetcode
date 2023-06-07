package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	cases := map[string]struct {
		expectedData string
		input        []string
	}{
		"Should return fl": {
			expectedData: "fl",
			input:        []string{"flower", "flow", "flight"},
		},
		"should return ": {
			expectedData: "",
			input:        []string{"dog", "racecar", "car"},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := longestCommonPrefixV2(cs.input)

			assert.Equal(t, cs.expectedData, resp)
		})
	}
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	input := []string{"flow", "falango", "falafel", "feitiço", "flito"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(input)
	}
}
func Benchmark_longestCommonPrefixV2(b *testing.B) {
	input := []string{"flow", "falango", "falafel", "feitiço", "flito"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefixV2(input)
	}
}
