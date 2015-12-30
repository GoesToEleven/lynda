package main

import "testing"

func BenchmarkResults(b *testing.B) {
	rows := getRows()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Results(rows)
	}
}
