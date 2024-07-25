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
func lengthOfLongestSubstringTwoDistinct(s string) int {
    freq := map[byte]int{}
    left := 0
    maxWin := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for len(freq) > 2 {
            leftChar := s[left]
            freq[leftChar]--
            if freq[leftChar] == 0 {
                delete(freq, leftChar)
            }
            left++
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}