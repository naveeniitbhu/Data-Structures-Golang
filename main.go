package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 4, 5, 6}, 7))
	fmt.Println("Reverse Int of ", reverseint(2356))
	fmt.Printf("Is Number %v Palindrome? : %v\n", 232, isPalindrome(232))
	fmt.Printf("Integere value of roman number %v is : %v \n", "LVIII", romanToInt("LVIII"))
	fmt.Printf("Integere value of roman number %v is : %v \n", "LVIII", romanToIntsolution2("LVIII"))
	fmt.Printf("Longest common preffix of %v is : %v \n", []string{"flower", "flow", "flight"}, longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("Is Valid Parenthese %v ? : %v\n", "({{)[]", isValid("({{)[]"))
	// fmt.Printf("Merge 2 sorted lists %v and %v ? : %v\n", []int{1, 2, 4}, []int{1, 2, 3}, mergeTwoLists([]int{1, 2, 4}, []int{1, 2, 3}))
	fmt.Println(fib(10))
	fibClosure()
	fmt.Println(fibRecursion(10))
	fmt.Printf("Removed duplicates from array %v - index-result %v\n", []int{0, 0, 1, 1, 1, 2, 2, 3, 4}, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 4}))
	fmt.Printf("Remove a particular elem from slice%v - index-result %v\n", []int{0, 0, 1, 1, 1, 2, 2, 3, 4}, removeElement([]int{0, 0, 1, 1, 1, 2, 2, 3, 4}, 2))
	fmt.Printf("Index of matching needle: %v in haystack: %v is: %v\n", "ll", "hello", haystackNeedle("hello", "ll"))
	fmt.Printf("Position of target:%v in array:%v is : %v\n", 5, []int{1, 3, 4, 6}, searchInsertPostionInArray([]int{1, 3, 4, 6}, 5))
	fmt.Printf("Number of Trailing zeroes in factorial of number: %v is %v\n", 25, trailingZeroes(25))
}
