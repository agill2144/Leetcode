func minWindow(s string, t string) string {
    freq := map[byte]int{}
    for i := 0; i < len(t); i++ {freq[t[i]]++}
    fullMatch := 0
    startIdx, endIdx := -1,-1
    left := 0
    for i := 0; i < len(s); i++ {
        val, ok := freq[s[i]]
        if ok {
            freq[s[i]]--
            if val == 1 {fullMatch++}
        }
        for left <= i && fullMatch == len(freq) {
            if startIdx == -1 || (i-left+1 < endIdx-startIdx+1) {
                startIdx = left
                endIdx = i
            }
            val, ok := freq[s[left]]
            if ok {
                freq[s[left]]++
                if val == 0 {
                    fullMatch--
                }
            }
            left++
        }
    }
    if startIdx == -1 {return ""}
    return s[startIdx:endIdx+1]
}