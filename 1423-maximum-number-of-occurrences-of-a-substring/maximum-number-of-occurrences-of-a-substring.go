// we are not count number of substrings that match the rule
// we are counting out of all substring that did match the rule, return the size
// of that 1 substring the occured the most
// { "aat": 11, "aa": 10 "bb": 10, "cba" : 92 }
// we return size of cba, since its the most frequently occuring substr that satifies all rules
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    freqMap := map[byte]int{}
    maxCount := 0
    left := 0
    wordSizeCounts := map[string]int{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        freqMap[char]++

        for len(freqMap) > maxLetters {
            leftChar := s[left]
            freqMap[leftChar]--
            if val := freqMap[leftChar]; val == 0 {
                delete(freqMap, leftChar)
            }
            left++
        }
        
        for i-left+1 >= minSize && i-left+1 <= maxSize {
            word := s[left:i+1]
            wordSizeCounts[word]++
            if val := wordSizeCounts[word]; val > maxCount {maxCount = val}

            leftChar := s[left]
            freqMap[leftChar]--
            if val := freqMap[leftChar]; val == 0 {
                delete(freqMap, leftChar)
            }
            left++            
        }
    }
    return maxCount
}