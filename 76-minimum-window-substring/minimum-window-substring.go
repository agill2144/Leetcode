func minWindow(s string, t string) string {
    freq := map[byte]int{}
    for i := 0; i < len(t); i++ {freq[t[i]]++}
    left := 0
    res := ""
    fullMatch := 0
    for i := 0; i < len(s); i++ {
        _, ok := freq[s[i]]
        if ok {
            freq[s[i]]--
            if freq[s[i]] == 0 {fullMatch++}
        }
        for fullMatch == len(freq) {
            subStr := s[left:i+1]
            if res == "" || len(subStr) < len(res) {
                res = subStr
            }
            _, ok = freq[s[left]]
            if ok {freq[s[left]]++}
            if freq[s[left]] == 1 {fullMatch--}
            left++
        }
    }
    return res
}