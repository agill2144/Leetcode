func removeKdigits(num string, k int) string {
    if k == len(num) {return "0"}
    
    st := []byte{}
    for i := 0; i < len(num); i++ {
        n := num[i]
        for len(st) != 0 && n < st[len(st)-1] && k != 0 {
            st = st[:len(st)-1]
            k--
        }
        // a number is not valid if it starts 0
        if len(st) == 0 && n == '0' {continue}
        st = append(st, n)
    }
    for k != 0 && len(st) != 0 {
        st = st[:len(st)-1]
        k--
    }

    // in case our stack is empty ( happens when 0 removed all elements from st and we do not push 0 into st)
    if len(st) == 0 {return "0"}

    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    return out.String()
}