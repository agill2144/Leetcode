// longest substr without repeating char
// substr = contagious part of string 
func lengthOfLongestSubstring(s string) int {
    left := 0
    maxWin := 0
    idx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        lastSeenAt, ok := idx[s[i]]
        if ok {
            if left <= lastSeenAt {
                left = lastSeenAt+1
            }
        }
        idx[s[i]] = i
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}