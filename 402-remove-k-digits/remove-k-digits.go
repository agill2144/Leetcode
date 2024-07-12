func removeKdigits(num string, k int) string {
    st := []byte{}
    for i := 0; i < len(num); i++ {
        for len(st) != 0 && num[i] < st[len(st)-1] && k > 0 {
            k--
            st = st[:len(st)-1]
        }
        if num[i] == '0' && len(st) == 0 {continue}
        st = append(st, num[i])
    }
    for k != 0 && len(st) != 0 {
        st = st[:len(st)-1]
        k--
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    if out.String() == "" {return "0"}
    return out.String()
}
