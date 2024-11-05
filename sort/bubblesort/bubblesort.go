package main

// import "fmt"

func BubbleSort(nums []int){
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swapped = true
			}
		}
	}
}

// func main() {
// 	nums := []int{78,231,1,23,14} 
// 	BubbleSort(nums)
// 	fmt.Println(nums) // 1,14,23,78,231
// }