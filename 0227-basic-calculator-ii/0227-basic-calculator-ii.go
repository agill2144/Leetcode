func calculate(s string) int {
    contr := 0
    calc := 0
    curr := 0
    var lastOp byte = '+'
    for i := 0; i < len(s); i++ {
        char := s[i]        
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if char == '+' || char == '-' || char == '*' || char == '/' || i == len(s)-1 {
            if lastOp == '+' {
                calc += curr
                contr = curr
            } else if lastOp == '-' {
                calc -= curr
                contr = -curr
            } else if lastOp == '*' {
                calc = calc-contr + contr * curr
                contr = contr * curr
            } else if lastOp == '/' {
                calc = calc-contr + contr / curr
                contr = contr / curr
            }
            curr = 0
            lastOp = char
        }
    }
    return calc
}