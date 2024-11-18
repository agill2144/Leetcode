func calculate(s string) int {
    st := []int{}
    var op byte = '+'
    curr := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if i == len(s)-1 || char == '+' || char == '-' || char == '*' || char == '/' {
            if op == '*' {
                st[len(st)-1] *= curr
            } else if op == '/' {
                st[len(st)-1] /= curr
            } else if op == '-' {
                st = append(st, -curr)
            } else {
                st = append(st, curr)
            }
            curr = 0
            op = char
        }
    }
    total := 0
    for len(st) != 0 {
        top := st[len(st)-1]; st = st[:len(st)-1]
        total += top
    }
    return total
}