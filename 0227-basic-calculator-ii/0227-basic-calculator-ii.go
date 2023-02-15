func calculate(s string) int {
    calc := 0
    curr := 0
    var sign byte = '+'
    tail := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if char == '+' || char == '-' || char == '*' || char == '/' || i == len(s)-1 {
            if sign == '+' {
                calc += curr
                tail = curr
            } else if sign == '-' {
                calc -= curr
                tail = -curr
            } else if sign == '*' {
                calc = calc - tail + (tail*curr)
                tail = tail * curr
            } else if sign == '/' {
                calc = calc - tail + (tail / curr)
                tail = tail / curr
            }
            curr = 0
            sign = char
        }
    }
    return calc
}