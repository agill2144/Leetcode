func removeKdigits(num string, k int) string {
    n := len(num)
    st := []byte{}
    for i := 0; i < n; i++ {
        curr := num[i]
        for len(st) != 0 && curr < st[len(st)-1] && k > 0 {
            st = st[:len(st)-1]
            k--
        }
        if len(st) == 0 && curr == '0' {continue}
        st = append(st, curr)
    }
    for k != 0 && len(st) > 0 {
        st = st[:len(st)-1]
        k--
    }
    if len(st) == 0 {return "0"}
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {out.WriteByte(st[i])}
    return out.String()
}