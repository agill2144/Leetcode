func calculate(s string) int {
    st := []int{}
    var lastOp byte = '+'
    curr := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if i == len(s)-1 || (char == '+' || char == '-' || char == '*' || char == '/') {
            if lastOp == '+' {
                st = append(st, curr)
            } else if lastOp == '-' {
                st = append(st, -curr)
            } else if lastOp == '*' {
                st[len(st)-1] *= curr
            } else if lastOp == '/' {
                st[len(st)-1] /= curr
            }
            curr = 0
            lastOp = char
        }
    }
    total := 0
    for len(st) != 0 {
        total += st[len(st)-1]
        st = st[:len(st)-1]
    }
    return total
}