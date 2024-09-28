func calculate(s string) int {
    st := []int{}
    curr := 0
    var lastOp byte = '+'
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if i == len(s)-1 || char == '*' || char == '/' || char == '-' || char == '+' {
            if lastOp == '*' {
                top := 1
                if len(st) > 0 {
                    top = st[len(st)-1]
                }
                top *= curr
                st[len(st)-1] = top
            } else if lastOp == '/' {
                top := 1
                if len(st) > 0 {
                    top = st[len(st)-1]
                }
                st[len(st)-1] = top/curr
            } else if lastOp == '+' {
                st = append(st, curr)
            } else if lastOp == '-' {
                st = append(st, -curr)
            }
            lastOp = char
            curr = 0
        }
    }
    total := 0
    for len(st) != 0 {
        total += st[len(st)-1]
        st = st[:len(st)-1]
    }
    return total
}