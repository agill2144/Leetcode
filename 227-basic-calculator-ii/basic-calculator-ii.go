func calculate(s string) int {
    res := 0
    curr := 0
    var lastOp byte = '+'
    prevContr := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if i == len(s)-1 || char == '*' || char == '-' || char == '+' || char == '/' {
            // we ran into a new operator
            // process lastOp 
            if lastOp == '+' {
                res += curr
                prevContr = +curr
            } else if lastOp == '-' {
                res -= curr
                prevContr = -curr
            } else if lastOp == '*' {
                // remove prevContr + process multiplication first ( i.e; prev number * curr )
                // and then add back to result
                // hence res - prevContr + (prevContr + curr)
                res = res - prevContr + (prevContr * curr)
                prevContr = prevContr*curr
            } else if lastOp == '/' {
                // same as multiplication
                res = res - prevContr + (prevContr / curr)
                prevContr = prevContr/curr
            }
            lastOp = char
            curr = 0
        }
    }
    return res
}