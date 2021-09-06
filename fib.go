package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	var sl []int
	sl = append(sl, 0)

	for i := 1; i < n; i++ {
		if i == 1 {
			sl = append(sl, 1)
		} else {
			sl = append(sl, sl[i-2]+sl[i-1])
		}
	}
	return sl[n-1]
}

// Using Closures

func fibClosure() {
	f := fibonacciCL()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacciCL() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}

}

// Using Recursion basic

func fibRecursion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibRecursion(n-1) + fibRecursion(n-2)
	}
}
