package main

import (
	"testing"
)

func BenchmarkTsuncon(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getContributions("diplozoon")
	}
}
