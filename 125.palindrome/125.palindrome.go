package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^A-Za-z0-9]`)

func isPalindromeV1(s string) bool {
	if len(s) < 2 {
		return true
	}

	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		if s[i] != s[(len(s)-1)-i] {
			return false
		}
	}
	return true
}

func isPalindromeV2(s string) bool {
	if len(s) < 2 {
		return true
	}

	s = strings.ToLower(s)
	fmt.Println(s)

	start, end := make([]byte, 0, len(s)), make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		init, finish := rune(s[i]), rune(s[(len(s)-1)-i])
		// Valida se o caracter isn't letter or number
		if unicode.IsLetter(init) || unicode.IsNumber(init) {
			start = append(start, s[i])
		}

		if unicode.IsLetter(finish) || unicode.IsNumber(finish) {
			end = append(end, s[(len(s)-1)-i])
		}
	}

	fmt.Println(fmt.Sprintf("%s", start))
	fmt.Println(fmt.Sprintf("%s", end))

	return bytes.NewBuffer(start).String() == bytes.NewBuffer(end).String()
}

func isPalindromeV3(s string) bool {
	if len(s) < 2 {
		return true
	}

	start, end := make([]byte, 0, len(s)), make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		init, finish := unicode.ToLower(rune(s[i])), unicode.ToLower(rune(s[(len(s)-1)-i]))
		// Valida se o caracter isn't letter or number
		if unicode.IsLetter(init) || unicode.IsNumber(init) {
			start = append(start, s[i])
		}

		if unicode.IsLetter(finish) || unicode.IsNumber(finish) {
			end = append(end, s[(len(s)-1)-i])
		}
	}

	return bytes.NewBuffer(start).String() == bytes.NewBuffer(end).String()
}

func main() {
	test := map[string]bool{
		"A man, a plan, a canal: Panama": true,
		"race a car":                     false,
		" ":                              true,
	}

	for value, valid := range test {
		result := isPalindromeV1(value)
		if result != valid {
			fmt.Printf("V1: value: %+v, expected: %+v, return: %+v\n", value, valid, result)
		}
	}
	for value, valid := range test {
		result := isPalindromeV2(value)
		if result != valid {
			fmt.Printf("V2: value: %+v, expected: %+v, return: %+v\n", value, valid, result)
		}
	}
}
