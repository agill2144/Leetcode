func minWindow(s string, t string) string {
    tMap := map[byte]int{}
    for i := 0; i < len(t); i++ {
        tMap[t[i]]++
    }
    fm := len(tMap)
    ans := ""
    left := 0
    for i := 0; i < len(s); i++ {
        _, ok := tMap[s[i]]
        if ok {
            tMap[s[i]]--
            if tMap[s[i]] == 0 {fm--}
            for fm == 0 {
                if ans == "" || i-left+1 < len(ans) {
                    ans = s[left:i+1]
                }
                _, ok2 := tMap[s[left]]
                if ok2 {
                    tMap[s[left]]++
                    if tMap[s[left]] == 1{fm++}
                }
                left++
            }
        }
    }
    return ans
}