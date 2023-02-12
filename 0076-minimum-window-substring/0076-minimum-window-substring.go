func minWindow(s string, t string) string {
    if len(t) > len(s) {return ""}
    
    freqMap := map[byte]int{}
    for i := 0; i < len(t); i++ {freqMap[t[i]]++}
    
    start := 0
    end := len(s)
    minSize := math.MaxInt64
    left := 0
    count := 0
    
    for i := 0; i < len(s); i++ {
        char := s[i]
        if _, ok := freqMap[char]; ok {
            freqMap[char]--
            if val := freqMap[char]; val == 0 {
                count++
            }
        }
        for count == len(freqMap) {
            subStrSize := i-left+1
            if subStrSize < minSize {
                minSize = subStrSize
                start = left
                end = i     
            }
            leftChar := s[left]
            if _, ok := freqMap[leftChar]; ok {
                freqMap[leftChar]++
                if val := freqMap[leftChar]; val == 1 {
                    count--
                }
            }
            left++
        }
    }
    if minSize == math.MaxInt64 {return ""}
    return string(s[start:end+1])
}