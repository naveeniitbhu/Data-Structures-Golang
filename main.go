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
}
