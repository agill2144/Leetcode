func findAnagrams(s string, p string) []int {
    pMap := map[byte]int{}
    for i := 0; i < len(p); i++ {
        pMap[p[i]]++
    }
    count := 0
    left := 0
    out := []int{}
    for i := 0; i < len(s); i++ {
        val, ok := pMap[s[i]]
        if ok {
            pMap[s[i]]--
            if val == 1 {
                count++
            }
        }
        
        if i-left+1 == len(p) {
            if count == len(pMap) {
                out = append(out, left)
            }
            leftChar := s[left]
            val, ok = pMap[leftChar]
            if ok {
                pMap[leftChar]++
                if val == 0 {count--}
            }
            left++
        }
    }
    return out
}