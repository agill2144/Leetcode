func minWindow(s string, t string) string {
    tFreq := map[byte]int{}
    for i := 0; i < len(t); i++ {tFreq[t[i]]++}
    fullMatch := 0
    left := 0
    start, end := -1,-1
    for i := 0; i < len(s); i++ {
        _, ok := tFreq[s[i]]
        if ok {
            tFreq[s[i]]--
            if tFreq[s[i]] == 0 {fullMatch++}
        }
        for fullMatch == len(tFreq) {
            if start == -1 || i-left+1 < end-start+1 {
                start = left
                end = i
            }
            if _, ok := tFreq[s[left]]; ok {
                tFreq[s[left]]++
                if tFreq[s[left]] ==  1 {fullMatch--}
            }
            left++
        }
    }
    if start == -1 {
        return ""
    }
    return s[start:end+1]
}