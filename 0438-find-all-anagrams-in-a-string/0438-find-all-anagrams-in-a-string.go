func findAnagrams(s string, p string) []int {
    pMap := map[byte]int{}
    for i := 0; i < len(p); i++ {
        pMap[p[i]]++
    }
    left := 0
    out := []int{}
    count := len(pMap)
    for i := 0; i < len(s); i++{
        char := s[i]
        if _, ok := pMap[char]; ok {
            pMap[char]--
            // we have a full match for this char ?
            // decrement its counter
            if pMap[char] == 0 {count--}
        }
        
        if i-left+1 == len(p) {
            // count will be 0 when all chars were fully matched 
            // with current window
            if count == 0 {
                out = append(out, left)
            }
            leftChar := s[left]
            if _, ok := pMap[leftChar]; ok {
                pMap[leftChar]++
                if pMap[leftChar] == 1 {count++}
            }
            left++
        }
    }
    return out
}