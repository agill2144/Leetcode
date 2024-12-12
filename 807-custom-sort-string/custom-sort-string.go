func customSortString(order string, s string) string {
    // sc = o(1)
    sFreq := map[byte]int{}
    // m = len(order)
    // n = len(s)
    // tc = o(n)    
    for i := 0; i < len(s); i++ {sFreq[s[i]]++}
    out := new(strings.Builder)
    // tc = o(n+m)
    for i := 0; i < len(order); i++ {
        count := sFreq[order[i]]
        for count != 0 { out.WriteByte(order[i]); count--}
        delete(sFreq, order[i])
    }
    // o(n)
    for i := 0; i < len(s); i++ {
        if sFreq[s[i]] > 0 {
            out.WriteByte(s[i])
        }
    }
    // total tc = o(n) + o(n+m) + o(n) = o(n+m) + (2n)
    // sc = o(1)
    return out.String()
}