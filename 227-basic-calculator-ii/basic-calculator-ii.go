func calculate(s string) int {
    curr := 0
    total := 0
    lastContr := 0
    var lastOp byte = '+'
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if i == len(s)-1 || (char == '+' || char == '-' || char == '*' || char == '/') {
            if lastOp == '+' {
                total += curr
                lastContr = curr
            } else if lastOp == '-' {
                total -= curr
                lastContr = -curr
            } else if lastOp == '*' {
                multiRes := lastContr * curr
                total = total - lastContr + multiRes
                lastContr = multiRes
            } else if lastOp == '/' {
                divRes := lastContr / curr
                total = total - lastContr + divRes
                lastContr = divRes
            }
            curr = 0
            lastOp = char
        }
    }
    return total
}