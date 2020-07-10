package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

	nums = []int{0, 4, 3, 0}
	target = 0
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))
}

// Runtime: 4 ms, faster than 93.93% of Go online submissions for Two Sum.
// Memory Usage: 3.8 MB, less than 35.30% of Go online submissions for Two Sum
// goos: darwin
// goarch: amd64
// Benchmark_twoSum-4    3882222       284 ns/op      32 B/op       2 allocs/op
func twoSum(nums []int, target int) []int {
	tmp := map[int]int{}
	for k, v := range nums {
		diff := target - v
		i, ok := tmp[diff]
		if ok {
			return []int{i, k}
		}

		tmp[v] = k
	}

	return []int{}
}
