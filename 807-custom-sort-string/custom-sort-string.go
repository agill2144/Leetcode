func customSortString(order string, s string) string {
    sFreq := map[byte]int{}
    for i := 0; i < len(s); i++ {sFreq[s[i]]++}
    out := new(strings.Builder)
    for i := 0; i < len(order); i++ {
        count := sFreq[order[i]]
        for count != 0 { out.WriteByte(order[i]); count--}
        delete(sFreq, order[i])
    }
    for i := 0; i < len(s); i++ {
        if sFreq[s[i]] > 0 {
            out.WriteByte(s[i])
        }
    }
    return out.String()
}