package main

import "fmt"

func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
	return nums
}

func main() {
	fmt.Println(SelectionSort([]int{5, 3, 2, 1, 0, 4}))
	// prints
	// [0, 1, 2, 3, 4, 5]
}