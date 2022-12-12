func removeKdigits(num string, k int) string {
    if len(num) == k {return "0"}
    st := []byte{}
    for i := 0; i < len(num); i++ {
        for len(st) != 0 && num[i] < st[len(st)-1] && k != 0 {
            st = st[:len(st)-1]
            k--
        }
        if len(st) == 0 && num[i] == '0' {continue}
        st = append(st, num[i])
    } 
    
    for k > 0 && len(st) != 0 {
        st = st[:len(st)-1]
        k--
    }
    if len(st) == 0 {return "0"}

    return string(st)
}