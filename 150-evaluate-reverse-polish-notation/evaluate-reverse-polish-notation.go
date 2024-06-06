func evalRPN(tokens []string) int {
    st := []int{}
    for i := 0; i < len(tokens); i++ {
        currToken := tokens[i]
        if currToken == "+" {
            second := st[len(st)-1]
            first := st[len(st)-2]
            st = st[:len(st)-2]
            st = append(st, first+second)
        } else if currToken == "-" {
            second := st[len(st)-1]
            first := st[len(st)-2]
            st = st[:len(st)-2]
            st = append(st, first-second)
        } else if currToken == "*" {
            second := st[len(st)-1]
            first := st[len(st)-2]
            st = st[:len(st)-2]
            st = append(st, first*second)
        } else if currToken == "/" {
            second := st[len(st)-1]
            first := st[len(st)-2]
            st = st[:len(st)-2]
            st = append(st, first/second)
        } else {
            curr := 0
            isNegative := false
            for k := 0; k < len(currToken); k++ {
                if k == 0 && currToken[k] == '-' {isNegative = true; continue}
                curr = curr * 10 + int(currToken[k]-'0')
            }
            if isNegative {
                curr *= -1
            }
            st = append(st, curr)
        }
    }
    return st[len(st)-1]
}