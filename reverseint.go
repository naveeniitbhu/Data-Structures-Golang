package main

import "math"

func reverseint(x int) int {
	sign := 1
	if x<0 {
		sign = -1
		x *= -1
	}
	res := 0
	for x>0 {
		res = res*10 + x%10
		x = x/10
		if res > int(math.Pow(2,31) - 1) {
			return 0
		}
	}

	if sign < 0 && res*sign < int(math.Pow(2,31)) {
		return 0
	}
	
	return res * sign
}