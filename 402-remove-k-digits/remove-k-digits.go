// next smaller on left (strictly smaller, not <= curr num, only < curr) 
func removeKdigits(num string, k int) string {
    if k == len(num) {return "0"}
    st := []byte{}
    for i := 0; i < len(num); i++ {
        curr := num[i]
        for len(st) != 0 && k > 0 && st[len(st)-1] > curr {
            st = st[:len(st)-1]
            k--
        }
        if len(st) == 0 && curr == '0' {continue}
        st = append(st, curr)
    }
    for (k != 0 && len(st) != 0) {
        st = st[:len(st)-1]
        k--
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        if i == 0 && st[i] == '0' {continue}
        out.WriteByte(st[i])
    }
    outStr := out.String()
    if outStr == "" {outStr = "0"}
    return outStr
}