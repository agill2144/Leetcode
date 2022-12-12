func firstUniqChar(s string) int {
    freqMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freqMap[s[i]]++
    }
    for i := 0; i < len(s); i++ {
        if val, _ := freqMap[s[i]]; val == 1 {
            return i
        }
    }
    return -1
}