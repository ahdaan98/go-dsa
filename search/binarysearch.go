package main

// import "fmt"

func BinarySearch(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return true
		}
	}

	return false
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(BinarySearch(nums, 1))
// }
