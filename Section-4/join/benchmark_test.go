package main

import "testing"

func BenchmarkJoinBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinBase("./t/r0.tbl", "./t/r2.tbl", []int{0}, []int{1})
		//JoinBase("./t/r2.tbl", "./t/r0.tbl", []int{0}, []int{1})
		//JoinBase("./t/r8.tbl", "./t/r5.tbl", []int{0}, []int{1})
		//JoinBase("./t/r5.tbl", "./t/r8.tbl", []int{0}, []int{1})
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join("./t/r0.tbl", "./t/r2.tbl", []int{0}, []int{1})
		//Join("./t/r2.tbl", "./t/r0.tbl", []int{0}, []int{1})
		//Join("./t/r8.tbl", "./t/r5.tbl", []int{0}, []int{1})
		//Join("./t/r5.tbl", "./t/r8.tbl", []int{0}, []int{1})
	}
}

func BenchmarkJoinExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinExample("./t/r0.tbl", "./t/r2.tbl", []int{0}, []int{1})
		//JoinExample("./t/r2.tbl", "./t/r0.tbl", []int{0}, []int{1})
		//JoinExample("./t/r8.tbl", "./t/r5.tbl", []int{0}, []int{1})
		//JoinExample("./t/r5.tbl", "./t/r8.tbl", []int{0}, []int{1})
	}
}
