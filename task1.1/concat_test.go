package main

import "testing"

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Plus("adsf", "afsdfsd")
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Builder("adsf", "afsdfsd")
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"adsf", "afsdfsd"}, " ")
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Buffer("adsf", "afsdfsd")
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprint("adsf", "afsdfsd")
	}
}
