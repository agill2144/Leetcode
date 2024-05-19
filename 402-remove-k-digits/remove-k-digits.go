func removeKdigits(num string, k int) string {
    st := []byte{} // chars
    n := len(num)
    for i := 0; i < n; i++ {
        curr := num[i]
        // am i(curr) your(top-of-st) nsr ?
        // if yes, pop it if we still have k
        for len(st) != 0 && curr < st[len(st)-1] && k != 0 {
            st = st[:len(st)-1]
            k--
        }

        if curr == '0' && len(st) == 0 {continue}
        st = append(st, curr)
    } 

    for k != 0 && len(st) != 0 {
        st = st[:len(st)-1]
        k--
    }
    if len(st) == 0 {return "0"}

    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    return out.String()
}

// 3456789