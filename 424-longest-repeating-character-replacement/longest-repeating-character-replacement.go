func characterReplacement(s string, k int) int {
    left := 0
    freq := map[byte]int{}
    majority := 0
    ans := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        majority = max(majority, freq[s[i]])
        size := i-left+1
        if size-majority <= k {
            ans = max(ans, size)
        } else {
            leftChar := s[left]
            freq[leftChar]--
            left++
        }
    }
    return ans
}