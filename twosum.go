package main


func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for index, num := range nums {
		targetSubstractbyNum := target - num
		if result, booleanCheck := myMap[targetSubstractbyNum]; booleanCheck {
			return []int{result, index}
		}
		myMap[num] = index
	}
	return []int{0, 0}
}

// Input
// [2,7,11,15]  9
// Output
// [0,1]