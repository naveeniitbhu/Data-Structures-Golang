package main

func searchInsertPostionInArray(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}

		if i == len(nums)-1 && nums[i] < target {
			return i + 1
		}
	}
	return 0
}
