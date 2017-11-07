package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{3, 2, 4}
	target := 6
	start := time.Now()
	ans := twoSum(nums, target)
	fmt.Println("excursion", time.Since(start))
	fmt.Println("Answer", ans)
}

func twoSum2(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for i, val := range nums {
		if res, ok := hashmap[val]; ok {
			return []int{res, i}
		}
		hashmap[target-val] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	result := []int{}
	ch := make(chan int)
	found := false
	for i, num := range nums {
		go func(nums []int, target, index, num int) {
			for j := index + 1; j < len(nums); j++ {
				if num+nums[j] == target {
					result = append(result, index, j)
					found = true
				}
			}
			ch <- 1
		}(nums, target, i, num)
	}

	for i := 0; i < len(nums); i++ {
		<-ch
	}

	if len(result) > 2 {
		result = result[:2]
	}

	return result
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i, num := range nums {
		diff := target - num
		if exist, j := contains(nums, diff); exist && i != j {
			result = append(result, i, j)
			break
		}
	}

	if len(result) > 2 {
		result = result[:2]
	}

	return result
}

func contains(s []int, e int) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}
