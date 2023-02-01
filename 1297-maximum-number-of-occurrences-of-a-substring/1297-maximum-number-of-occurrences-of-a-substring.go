func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    if maxSize > len(s) || maxSize < minSize || len(s) == 0 {return 0}

    wordMap := map[string]int{}
    maxCount := 0
    windowMap := map[string]int{}    
    left := 0
    
    for right := 0; right < len(s); right++ {
        strChar := string(s[right])
        windowMap[strChar]++
        
        for len(windowMap) > maxLetters {
            leftStrChar := string(s[left])
            windowMap[leftStrChar]--
            if val := windowMap[leftStrChar]; val == 0 {
                delete(windowMap, leftStrChar)
            }
            left++
        }
        
        windowStr := s[left:right+1]
        for len(windowStr) >= minSize && len(windowStr) <= maxSize {
            wordMap[windowStr]++
            if val := wordMap[windowStr]; val > maxCount {maxCount = val}
            leftStrChar := string(s[left])
            windowMap[leftStrChar]--
            if val := windowMap[leftStrChar]; val == 0 {
                delete(windowMap, leftStrChar)
            }
            left++
            windowStr = s[left:right+1]
        }
    }
    
    fmt.Println(wordMap, "")
    return maxCount
}