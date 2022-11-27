package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func FuzzMergesort(f *testing.F) {
	fuzzSort(f, mergesort)
}

func FuzzConcsort(f *testing.F) {
	fuzzSort(f, Concsort)
}

func fuzzSort(f *testing.F, sort func([]int) []int) {
	f.Fuzz(func(t *testing.T, seed int64, L int) {
		L = abs(L)
		rand.Seed(seed)
		arr := rand.Perm(L)
		sorted := sort(arr)
		assert.Len(t, sorted, L)
		for i := range sorted {
			assert.Equal(t, sorted[i], i)
		}
	})
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// func FuzzMergesort(f *testing.F) {
// 	// can't fuzz array
// 	f.Fuzz(func(t *testing.T, arr []int) {
// 		sorted := mergesort(arr)
// 		assert.Len(t, sorted, len(arr))
// 		for i := 0; i+1 < len(sorted); i++ {
// 			assert.LessOrEqual(t, sorted[i], sorted[i+1])
// 		}
// 	})
// }

func BenchmarkMergsort100(b *testing.B) {
	benchmarkSort(b, mergesort, 100)
}

func BenchmarkConcsort100(b *testing.B) {
	benchmarkSort(b, Concsort, 100)
}

func BenchmarkMergsort1000(b *testing.B) {
	benchmarkSort(b, mergesort, 1000)
}

func BenchmarkConcsort1000(b *testing.B) {
	benchmarkSort(b, Concsort, 1000)
}

func BenchmarkMergsort10000(b *testing.B) {
	benchmarkSort(b, mergesort, 10000)
}

func BenchmarkConcsort10000(b *testing.B) {
	benchmarkSort(b, Concsort, 10000)
}

func BenchmarkMergsort100000(b *testing.B) {
	benchmarkSort(b, mergesort, 100000)
}

func BenchmarkConcsort100000(b *testing.B) {
	benchmarkSort(b, Concsort, 100000)
}

func benchmarkSort(b *testing.B, sort func(arr []int) []int, L int) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		rand.Seed(time.Now().Unix())
		arr := rand.Perm(L)
		b.StartTimer()
		sort(arr)
	}
}
