func minSteps(s string, t string) int {
    sMap := map[byte]int{}
    for i := 0; i < len(s); i++ {sMap[s[i]]++}
    count := 0
    for i := 0; i < len(t); i++ {
        char := t[i]
        val, ok := sMap[char]
        if ok {
            sMap[char]--
            if val == 1 {
                delete(sMap, char)
            }
        } else {
            count++
        }
    }
    return count
}