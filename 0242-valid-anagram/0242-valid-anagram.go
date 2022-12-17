func isAnagram(s string, t string) bool {
    if len(s) != len(t) {return false}
    sFreqMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        sFreqMap[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        char := t[i]
        count, ok := sFreqMap[char]
        if !ok || count == 0 {
            return false
        }
        sFreqMap[char]--
    }
    return true
    
}