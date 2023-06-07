package main

import "testing"

func Benchmark_isPalindromeV1(b *testing.B) {
	palindrome := "A man, a plan, a canal: Panama"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		isPalindromeV1(palindrome)
	}
}

func Benchmark_isPalindromeV2(b *testing.B) {
	palindrome := "A man, a plan, a canal: Panama"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		isPalindromeV2(palindrome)
	}
}
