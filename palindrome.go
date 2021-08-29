package main

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    if x < 10 {
        return true
    }

    res, xo := 0, x
    
    for xo != 0 {
        pop := xo%10
        xo /= 10
        res = res*10 + pop
    }
    
    return res == x
}