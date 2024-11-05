package main

import "fmt"

func InsertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	return nums
}

func main() {
	fmt.Println(InsertionSort([]int{5, 3, 2, 1, 0, 4}))
	// prints
	// [0, 1, 2, 3, 4, 5]
}