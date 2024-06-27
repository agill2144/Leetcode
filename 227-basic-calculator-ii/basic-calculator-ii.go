func calculate(s string) int {
    st := []int{}
    digit := 0
    var operator byte = '+'
    for i := 0; i < len(s); i++ {
        char := s[i]
        if s[i] >= '0' && s[i] <= '9' {
            charInt := int(char-'0')
            digit = digit * 10 + charInt
        }
        if char == '+' || char == '-' || char == '*' || char == '/' || i == len(s)-1 {
            if operator == '+' {
                st = append(st, digit)
            } else if operator == '-' {
                st = append(st, -digit)
            } else if operator == '*' {
                top := st[len(st)-1]
                res := top*digit
                st[len(st)-1] = res
            } else if operator == '/' {
                top := st[len(st)-1]
                res := top/digit
                st[len(st)-1] = res
            }
            digit = 0
            operator = char
        }
    }
    res := 0
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        res += top
    }
    return res
}