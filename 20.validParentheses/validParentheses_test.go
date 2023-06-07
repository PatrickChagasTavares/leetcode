package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	cases := map[string]struct {
		expectedData bool
		input        string
	}{
		// "Should return ()=true": {
		// 	expectedData: true,
		// 	input:        "()",
		// },
		"Should return ()[]{}=true": {
			expectedData: true,
			input:        "()[]{}",
		},
		// "Should return {()[]}=false": {
		// 	expectedData: false,
		// 	input:        "{()[]}",
		// },
		// "Should return {}[)=false": {
		// 	expectedData: false,
		// 	input:        "{}[)",
		// },
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := isValid(cs.input)

			assert.Equal(t, cs.expectedData, resp)
		})
	}
}

func Test_isValidV2(t *testing.T) {
	cases := map[string]struct {
		expectedData bool
		input        string
	}{
		"Should return ()=true": {
			expectedData: true,
			input:        "()",
		},
		"Should return {()[]}=true": {
			expectedData: true,
			input:        "{()[]}",
		},
		"Should return {}[)=false": {
			expectedData: false,
			input:        "{}[)",
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := isValidV2(cs.input)

			assert.Equal(t, cs.expectedData, resp)
		})
	}

}
func Test_isValidV3(t *testing.T) {
	cases := map[string]struct {
		expectedData bool
		input        string
	}{
		// "Should return ()=true": {
		// 	expectedData: true,
		// 	input:        "()",
		// },
		"Should return ()[]{}=true": {
			expectedData: true,
			input:        "()[]{}",
		},
		// "Should return (]=false": {
		// 	expectedData: false,
		// 	input:        "(]",
		// },
		// "Should return ([)]=false": {
		// 	expectedData: false,
		// 	input:        "([)]",
		// },
		// "Should return {()[]}=false": {
		// 	expectedData: false,
		// 	input:        "{()[]}",
		// },
		// "Should return {}[)=false": {
		// 	expectedData: false,
		// 	input:        "{}[)",
		// },
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := isValidV3(cs.input)

			assert.Equal(t, cs.expectedData, resp)
		})
	}
}
