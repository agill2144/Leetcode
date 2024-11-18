func minRemoveToMakeValid(s string) string {
    st := []int{} // idx of open parans
    invalid := map[int]bool{} // idxs of all invalid chars that should not be part of final string
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            st = append(st, i)
        } else if s[i] == ')' {
            if len(st) == 0 {
                invalid[i] = true
            } else {
                st = st[:len(st)-1]
            }
        }
    }
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        invalid[top] = true
    }
    res := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if invalid[i] {continue}
        res.WriteByte(s[i])
    }
    return res.String()
}