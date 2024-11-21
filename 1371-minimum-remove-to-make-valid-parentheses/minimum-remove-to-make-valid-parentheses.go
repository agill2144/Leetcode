func minRemoveToMakeValid(s string) string {
    invalidIdxs := map[int]bool{}
    st := []int{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            st = append(st, i)
        } else if s[i] == ')' {
            if len(st) == 0 {
                invalidIdxs[i] = true
                continue
            }
            st = st[:len(st)-1]
        }
    }
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        invalidIdxs[top] = true 
    }
    res := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if invalidIdxs[i] {continue}
        res.WriteByte(s[i])
    }
    return res.String()
}