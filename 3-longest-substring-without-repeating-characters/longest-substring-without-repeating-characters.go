// longest substr without repeating char
// substr = contagious part of string 
func lengthOfLongestSubstring(s string) int {
    left := 0
    maxWin := 0
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        _, ok := freq[s[i]]
        for ok {
            freq[s[left]]--
            if freq[s[left]] == 0 {
                delete(freq, s[left])
            }
            left++
            _, ok = freq[s[i]]
        }
        freq[s[i]]++
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}