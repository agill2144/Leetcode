/*
substring == sliding window
instead of sliding the window 1-by-1 from left, make the left ptr take 1 big jump
*/
func lengthOfLongestSubstring(s string) int {
    idx := map[byte]int{}
    maxWin := 0
    left := 0
    for i := 0; i < len(s); i++ {
        val, ok := idx[s[i]]
        if ok && left <= val {left = val+1}
        idx[s[i]] = i
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}