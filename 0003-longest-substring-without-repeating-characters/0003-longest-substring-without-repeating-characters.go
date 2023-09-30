func lengthOfLongestSubstring(s string) int {
    windowSize := 0
    left := 0
    charIdx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        lastSeenAtIdx, ok := charIdx[char]
        if ok && left <= lastSeenAtIdx {
            left = lastSeenAtIdx+1
        }
        charIdx[char]=i
        if i-left+1 > windowSize {
            windowSize = i-left+1
        }
    }
    return windowSize
}
/*
    naive sliding window
    - try to insert a char in window 
    - while we cannot insert this char ( because this is a dupe )
        - start removing chars from the left side
    - once we have successfully inserted a char
    - check size of window and save if its better
    
    time = o(s)
    space = o(26) = o(1)
*/
// func lengthOfLongestSubstring(s string) int {
//     left := 0
//     windowState := map[byte]struct{}{}
//     windowSize := 0
//     for i := 0; i < len(s); i++ {
//         char := s[i]
//         _, ok := windowState[char]
//         for ok {
//             leftChar := s[left]
//             delete(windowState, leftChar)
//             left++
//             _, ok = windowState[char]
//         }
//         windowState[char] = struct{}{}
//         if i-left+1 > windowSize {
//             windowSize = i-left+1
//         }
//     }
//     return windowSize
// }