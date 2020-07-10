package main

import (
	"testing"
)

func Benchmark_twoSum(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 91,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				nums:   []int{0, 3, 4, 0},
				target: 0,
			},
			want: []int{0, 3},
		},
	}

	for i := 0; i < b.N; i++ {
		for no, t := range tests {
			take := twoSum(t.args.nums, t.args.target)
			for k, v := range take {
				if v != t.want[k] {
					b.Fatalf("No.%d 不一樣 %v != %v", no+1, take, t.want)
					break
				}
			}
		}
	}
}
