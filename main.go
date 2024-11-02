package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if idx, found := indexMap[diff]; found {
			return []int{idx, i}
		}

		indexMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)

	if result != nil {
		fmt.Printf("Indices: %v\n", result)
	} else {
		fmt.Println("No two sum solution found.")
	}
}
