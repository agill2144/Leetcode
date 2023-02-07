func customSortString(order string, s string) string {
    sFreqMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        sFreqMap[s[i]]++
    }
    out := new(strings.Builder)
    for i := 0; i < len(order); i++ {
        count, ok := sFreqMap[order[i]]
        if ok {
            for k := 0; k < count; k++ {
                out.WriteByte(order[i])
            }
        }
        delete(sFreqMap, order[i])
    }
    
    for i := 0; i < len(s); i++ {
        if _, ok := sFreqMap[s[i]]; ok {
            out.WriteByte(s[i])
        }
    }
    return out.String()
}