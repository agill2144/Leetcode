func findAnagrams(s string, p string) []int {
    if len(p) > len(s) {return nil}
    
    pMap := map[byte]int{}
    for i := 0; i < len(p); i++ {pMap[p[i]]++}
    count := len(pMap)
    out := []int{}
    
    left := 0
    for right := 0; right < len(s); right++ {
        char := s[right]
        val, ok := pMap[char]
        if ok {
            pMap[char]--
            if val == 1 {count--} 
        }
        if right-left+1 == len(p) {
            if count == 0 {out = append(out, left)}
            leftChar := s[left]
            if _, ok := pMap[leftChar]; ok {
                pMap[leftChar]++
                if val := pMap[leftChar]; val == 1 {count++}
            }
            left++
        }
    }
    return out
    
}