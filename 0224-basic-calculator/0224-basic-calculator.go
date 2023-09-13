type stackNode struct {
    sign byte
    n int
}
func calculate(s string) int {
    st := []stackNode{}
    curr := 0
    calc := 0
    s = strings.TrimSpace(s)
    var sign byte = '+'
    for i := 0; i < len(s); i++ {
        char := s[i]
        if (char >= '0' && char <= '9') {
            curr = curr * 10 + int(char-'0')
            if i == len(s)-1 {
                if sign == '+' {calc += curr } else { calc -= curr }
            }
        } else if char == '+' || char == '-' {
            // solve curr
            if sign == '+' {
                calc += curr
            } else {
                calc -= curr
            }
            // reset curr to 0
            curr = 0
            // set the sign
            sign = char
        } else if char == '(' {
            // push parent to stack; which will be used later for a child when the inner child runs into a closing bracket
            st = append(st, stackNode{sign: sign, n: calc})
            // fmt.Println("parent num: ", calc)
            // fmt.Println("parent sign: ", string(sign))
            // reset
            calc = 0
            sign = '+'
        } else if char == ')' {
            // solve curr expression
            if sign == '+' {
                calc += curr
            } else {
                calc -= curr
            }
            // reset
            curr = 0
            sign = '+'
            // combine with parent
            parent := st[len(st)-1]; st = st[:len(st)-1]
            if parent.sign == '+' {
                parent.n += calc
            } else if parent.sign == '-' {
                parent.n -= calc
            }
            calc = parent.n
        }
    }
    return calc
}