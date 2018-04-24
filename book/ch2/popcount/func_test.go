package popcount_test

import "testing"
import "book/ch2/popcount"

func BenchmarkPopCount8Loop(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount8Loop(uint64(i))
    }
}

func BenchmarkPopCount64Loop(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount64Loop(uint64(i))
    }
}

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount(uint64(i))
    }
}

func BenchmarkPopCountMagic(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCountMagic(uint64(i))
    }
}
