func minWindow(s string, t string) string {
    freq := map[byte]int{}
    for i := 0; i < len(t); i++ {freq[t[i]]++}
    fullMatch := 0
    left := 0
    res := ""
    for i := 0; i < len(s); i++ {
        if _, ok := freq[s[i]]; ok {
            freq[s[i]]--
            if freq[s[i]] == 0 {fullMatch++}
            for left <= i && fullMatch == len(freq) {
                if res == "" || i-left+1 < len(res) {
                    res = s[left:i+1]
                }
                if _, ok := freq[s[left]]; ok {
                    freq[s[left]]++
                    if freq[s[left]] == 1 {
                        fullMatch--
                    }
                }
                left++
            }
        }
    }
    return res
}