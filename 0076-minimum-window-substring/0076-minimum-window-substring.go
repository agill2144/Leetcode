func minWindow(s string, t string) string {
    tMap := map[byte]int{}
    for i := 0; i < len(t); i++ { tMap[t[i]]++ }
    ans := ""
    count := 0
    left := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        val, ok := tMap[char]
        if ok {
            tMap[char]--
            if val == 1 {
                count++
            }
        }
        for count == len(tMap) {
            size := i-left+1
            if ans == "" || size < len(ans) {
                ans = s[left:i+1]
            }
            leftChar := s[left]
            val , ok := tMap[leftChar]
            if ok {
                tMap[leftChar]++
                if val == 0 {
                    count--
                }
            }
            left++
        }
    }
    return ans
}