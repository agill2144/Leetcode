func evalRPN(tokens []string) int {
    st := []int{}
    for i := 0; i < len(tokens); i++ {
        token := tokens[i]
        n, err := strconv.Atoi(token)
        if err == nil {
            st = append(st, n)
            continue
        } 
        first := st[len(st)-1]
        second := st[len(st)-2]
        st = st[:len(st)-2]
        if token == "+" {st = append(st, first+second)}
        if token == "-" {st = append(st, second-first)}
        if token == "*" {st = append(st, first*second)}
        if token == "/" {st = append(st, second/first)}        
    }
    return st[len(st)-1]
}