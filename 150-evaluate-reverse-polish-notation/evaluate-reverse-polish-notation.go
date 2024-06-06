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
            nInt, _ := strconv.Atoi(currToken)
            st = append(st, nInt)
        }
    }
    return st[len(st)-1]
}