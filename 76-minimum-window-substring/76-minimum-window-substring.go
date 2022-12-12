
/*
    sliding window + a freqMap for the pattern to look for in a window
    - keep in mind, we are looking for the smallest possible window
    
    time: o(s) + o(t)
    space: o(t) -- freqMap
*/
func minWindow(s string, t string) string {
    tMap := map[string]int{}
    for _, ele := range t {
        tMap[string(ele)]++
    }
    
    windowStart := 0
    count := 0
    startIdx := 0
    endIdx := 0
    minWinSize := len(s)+1
    
    for windowEnd := 0; windowEnd < len(s); windowEnd++ {
        char := string(s[windowEnd])
        _, ok := tMap[char]
        if ok {
            tMap[char]--
            if val := tMap[char]; val == 0{
                count++
            }
        }
        for count == len(tMap) {
            cl := windowEnd - windowStart + 1
            if cl < minWinSize {
                minWinSize = cl
                startIdx = windowStart
                endIdx = windowEnd
            }
            
            leftChar := string(s[windowStart])
            val, ok := tMap[leftChar]
            if ok {
                if val == 0 {
                    count--   
                }
                tMap[leftChar]++
            }
            windowStart++
            
        }
    }
    if minWinSize == len(s)+1 {
        return ""
    }
    return string(s[startIdx:endIdx+1])
}