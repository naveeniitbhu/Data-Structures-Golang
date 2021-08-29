package main

func romanToInt(s string) int {
    m := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    result := m[s[len(s)-1]]
    for i:=len(s)-2; i>=0; i-- {
        if m[s[i]] < m[s[i+1]] {
            result -= m[s[i]]
        } else {
            result += m[s[i]]
        }
    }
    return result 
    
}

func romanToIntsolution2(s string) int {
    m := make(map[byte]int)

    m['I'] = 1
    m['V'] = 5
    m['X'] = 10
    m['L'] = 50
    m['C'] = 100
    m['D'] = 500
    m['M'] = 1000
    
    res := 0
    
    for i:=0; i<len(s);i++ {
        if i<len(s)-1 {
            if s[i]=='I'  && (s[i+1]=='V' || s[i+1]=='X') {
                res = res - 1
            } else if s[i]=='X' && (s[i+1]=='L' || s[i+1]=='C') {
                res = res - 10
            } else if s[i]=='C' && (s[i+1]=='D' || s[i+1]=='M') {
                res = res - 100
            } else {
                res = res + m[s[i]]
            }
        } else {
            res = res + m[s[i]]            
        }


    }
    return res
}