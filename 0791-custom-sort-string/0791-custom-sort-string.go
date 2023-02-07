/*
    s must be in the same form as order
    - freqMap of s chars
    - loop over chars in order string
        - if this char exists in sFreqMap
        - repeat this char that many times in the output string
        - since we have already used this char from order string, we can delete this char from the sFreqMap
    - Finally append whatever else is remaining from s ( compare with sFreqMap )
    
    time: o(len(s)) + o(len(order))
    space: o(len(s))
*/
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
            // o(s) ; worst case, all chars == this char in order
            // order = "b"
            // s = "bbbbbbbbbbb"
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
    
    // total time : o(s)+o(order*maxRepeatsOfACharInS)
    // total space: o(1)
    return out.String()
}