/*
    sliding window nested iteration optimizations;
    - dont shrink the window , only increase it
    - maintain a idx map
    - add each ith char in our idx map
    - if we have a valid window; len(idx) <= 2 
        - take win size and save its better than last
    - if we have a invalid window; len(idx) > 2
        - then shrink window while keeping the same size
        - if the left char idx in our idx-map comes before or at the position as left ptr
            - this char is being removed from our window
            - therefore we can remove its idx from our idx-map
            - because the position we are tracking is leaving our window
    time = o(n)
    space = o(1)

*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
    maxWin := 0
    left := 0
    idx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        idx[s[i]] = i
        if len(idx) <= 2 {
            maxWin = max(maxWin, i-left+1)
        } else {
            leftChar := s[left]
            if idx[leftChar] == left {
                delete(idx, leftChar)
            }
            left++
        }
    }
    return maxWin
}
/*
    sliding window nested iteration optimizations;
    - can we take a big jump at once?
    - yes
    - maintain a idx map
    - if the new char being inserted increases the size of idx map to 3
    - we know we need to escape the earliest idx
    - therefore find the earliest idx from our idx map
    - make left ptr take a bigger jump
    - delete the key thats leaving our window
    - now insert the new char in idx map
    - now we have a valid window

    time = o(n)
    space = o(1)

*/
// func lengthOfLongestSubstringTwoDistinct(s string) int {
//     idx := map[byte]int{}
//     maxWin := 0
//     left := 0
//     for i := 0; i < len(s); i++ {
//         _, ok := idx[s[i]]
//         if !ok && len(idx) == 2 { // we are adding a 3rd char now
//             // we need to escape the earliest possible idx
//             minIdx := len(s)+1
//             var minIdxChar byte
//             for k, v := range idx {
//                 if v < minIdx {
//                     minIdx = v
//                     minIdxChar = k
//                 }
//             }
//             if left <= minIdx {
//                 left = minIdx+1
//                 delete(idx, minIdxChar)
//             }
//         }
//         idx[s[i]] = i
//         maxWin = max(maxWin, i-left+1)
//     }
//     return maxWin
// }

/*
    naive slinding window
    - maintain a left and right ptr
    - maintain a freq map
    - as soon as our freq map size exceeds, shrink window and make it valid
    - shrinking happens 1 by 1 from left side, 1 idx at a time
    - and we keep updating our freq
    
    time = o(2n) = o(n)
    space = o(1)
*/
// func lengthOfLongestSubstringTwoDistinct(s string) int {
//     freq := map[byte]int{}
//     left := 0
//     maxWin := 0
//     for i := 0; i < len(s); i++ {
//         freq[s[i]]++
//         for len(freq) > 2 {
//             leftChar := s[left]
//             freq[leftChar]--
//             if freq[leftChar] == 0 {
//                 delete(freq, leftChar)
//             }
//             left++
//         }
//         maxWin = max(maxWin, i-left+1)
//     }
//     return maxWin
// }