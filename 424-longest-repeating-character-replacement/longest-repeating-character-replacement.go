func characterReplacement(s string, k int) int {
    left := 0
    freq := map[byte]int{}
    majority := 0
    ans := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        majority = max(majority, freq[s[i]])
        size := i-left+1
        for size-majority > k {
            leftChar := s[left]
            freq[leftChar]--
            if freq[leftChar] == 0 {delete(freq, leftChar)}
            left++
            size = i-left+1
            majority = -1
            for _, v := range freq {majority = max(majority, v)}
        }

        ans = max(ans, size)
    }
    return ans
}