package main

// import "fmt"

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// func main() {
// 	nums := []int{2,7,11,15}
// 	target := 9
// 	fmt.Printf("Input: %v\ntarget = %v\n",nums,target)
// 	fmt.Printf("Output: %v",TwoSum(nums,target))
// }

// func TwoSum(nums []int, target int) []int {
// 	mp := make(map[int]int)

// 	for ind, n := range nums {
// 		ans := target - n
// 		if _,exist := mp[ans]; exist{
// 			return []int{mp[ans],ind}
// 		}

// 		mp[n] = ind
// 	}

// 	return nil
// }