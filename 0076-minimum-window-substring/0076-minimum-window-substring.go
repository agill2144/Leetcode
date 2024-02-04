func minWindow(s string, t string) string {
    tMap := map[byte]int{}
    for i := 0; i < len(t); i++ { tMap[t[i]]++ }
    start := 0
    end := 0
    minSize := math.MaxInt64
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
            if size < minSize {
                minSize = size
                start = left
                end = i
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
    if minSize == math.MaxInt64 {return ""}
    
    return s[start:end+1]
}