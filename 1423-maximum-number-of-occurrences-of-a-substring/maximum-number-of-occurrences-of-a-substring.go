func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    left := 0
    maxSubStrFreq := 0
    subStrFreq := map[string]int{}
    charFreq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        charFreq[s[i]]++
        if len(charFreq) > maxLetters {
            leftChar := s[left]
            charFreq[leftChar]--
            if charFreq[leftChar] == 0 {delete(charFreq, leftChar)}
            left++
        } else {
            for i-left+1 >= minSize && i-left+1 <= maxSize {
                subStr := s[left:i+1]
                subStrFreq[subStr]++
                maxSubStrFreq = max(maxSubStrFreq, subStrFreq[subStr])
                leftChar := s[left]
                charFreq[leftChar]--
                if charFreq[leftChar] == 0 {delete(charFreq, leftChar)}
                left++
            }
        }
    }
    return maxSubStrFreq
}

/*
    classic sliding window
    - make sure the window is valid first!
    - then form all substrings within the valid window
        - take the first string of the window
        - now shrink the window from left side
        - again check if its valid ( winSize >= minSize and winSize <= maxSize )
        - form the next string in the window
        - then repeat..
*/
// func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
//     left := 0
//     maxSubStrFreq := 0
//     subStrFreq := map[string]int{}
//     charFreq := map[byte]int{}
//     for i := 0; i < len(s); i++ {
//         charFreq[s[i]]++
//         for len(charFreq) > maxLetters {
//             leftChar := s[left]
//             charFreq[leftChar]--
//             if charFreq[leftChar] == 0 {delete(charFreq, leftChar)}
//             left++
//         }
        
//         for i-left+1 >= minSize && i-left+1 <= maxSize {
//             subStr := s[left:i+1]
//             subStrFreq[subStr]++
//             maxSubStrFreq = max(maxSubStrFreq, subStrFreq[subStr])
//             leftChar := s[left]
//             charFreq[leftChar]--
//             if charFreq[leftChar] == 0 {delete(charFreq, leftChar)}
//             left++
//         }
//     }
//     return maxSubStrFreq
// }
