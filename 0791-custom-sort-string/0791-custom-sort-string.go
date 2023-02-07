func customSortString(order string, s string) string {
    sFreqMap := map[byte]int{} // space: o(1), there are constant number of characters in this world.
    
    // time: o(s)
    for i := 0; i < len(s); i++ {
        sFreqMap[s[i]]++
    }
    
    out := new(strings.Builder)
    
    
    // time: o(order)
    for i := 0; i < len(order); i++ {
        count, ok := sFreqMap[order[i]]
        if ok {
            for k := 0; k < count; k++ {
                out.WriteByte(order[i])
            }
        }
        delete(sFreqMap, order[i])
    }
    
    // time: o(s)
    for i := 0; i < len(s); i++ {
        if _, ok := sFreqMap[s[i]]; ok {
            out.WriteByte(s[i])
        }
    }
    
    // total time : o(2s)+o(order)
    // total space: o(1)
    return out.String()
}