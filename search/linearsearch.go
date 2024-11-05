package main

// import "fmt"

func LinearSearch(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7}
// 	fmt.Println(LinearSearch(nums,4))
// }