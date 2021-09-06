package main

func removeElement(nums []int, val int) int {
	curr := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[curr] = nums[i]
			curr++
		}
	}
	return curr
}
