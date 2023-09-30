/*
    classic sliding window - "find the length of the longest substring"
    sliding window type = Dynamic ( we need the longest one possible )
    
    approach:
    - instead of trying to slide the window from left side until the $this char is no longer part of our window
    - we can just track where $this char was last seen at position in a map instead of set
    
    - if we are trying to insert a char
    - and this char already exists in our window
    - we cannot insert this char until the same char is no longer part of our window
    - instead of slowly shrinking the left side 1 char at a time,
    - if we kept track of where $this char exists, we can fast-forward our window sliding from left side
    - eg window : {a:0, b:1, c:2, d:3}
    - leftPtr = 0
    - curr char to instert is "d"
    - we cannot insert $currChar until "d" is no longer part of the window
    - instead of moving leftPtr 1 by 1, we can tell leftPtr to go to idx 4 ( window["d"]+1 )
    - this means, that char "d" is now no longer part of our window (i.e it does not exists within leftPtr and i ptr)
    - and now we can insert this char in our window at ith position 
    - windowState will become = {a:0, b:1, c:2, d:4}
        - windowState is no longer in sync with the chars between leftPtr and i and THATS OKAY!
    
    time: o(n)
    space: o(n) 
*/
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