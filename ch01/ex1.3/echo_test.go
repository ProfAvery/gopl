package main

import (
    "testing"
)

var args = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}

func BenchmarkEcho1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Echo1(args)
    }
}

func BenchmarkEcho2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Echo2(args)
    }
}

func BenchmarkEcho3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Echo3(args)
    }
}

