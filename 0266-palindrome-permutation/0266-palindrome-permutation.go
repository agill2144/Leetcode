func canPermutePalindrome(s string) bool {
    freqMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freqMap[s[i]]++
    }
    for k, v := range freqMap {
        if v % 2 == 0 {
            delete(freqMap, k)
        } else {
            // v = 13
            freqMap[k] = v-(v-1)
        }
    }
    if len(freqMap) <= 1{return true}
    return false
}