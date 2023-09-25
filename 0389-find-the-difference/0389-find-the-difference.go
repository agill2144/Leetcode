func findTheDifference(s string, t string) byte {
    freqMap := map[byte]int{}
    for i := 0 ; i < len(s); i++ {
        freqMap[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        freqMap[t[i]]--
    }
    var out byte
    for k, v := range freqMap {
        if v != 0 {out = k; break}
    }
    return out
}