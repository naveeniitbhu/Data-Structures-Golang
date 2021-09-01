package main

func isValid(s string) bool {
    var stack []rune = make([]rune, len(s))
    m := map[rune]rune{
        ')':'(',
        '}':'{',
        ']':'[',
    }
    
    top := 0
    
    for _, char := range s {
        if char=='(' || char=='[' || char=='{' {
            stack[top]=char
            top++
        } else {
            if top==0 || m[char] != stack[top-1] {
                return false
            }
            top--
        }
    }
    if top>0 {
        return false
    }
    return true
}