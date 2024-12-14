func minRemoveToMakeValid(s string) string {
    st := []int{} // idx
    invalid := map[int]bool{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            st = append(st, i)
        } else if s[i] == ')' {
            if len(st) == 0 {invalid[i] = true; continue}
            st = st[:len(st)-1]
        }
    }
    for i := 0; i < len(st); i++ {invalid[st[i]] = true}
    out := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if invalid[i] {continue}
        out.WriteByte(s[i])
    }
    return out.String()
}