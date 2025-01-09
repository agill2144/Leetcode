func lengthOfLongestSubstring(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        set := make([]bool, 256)
        for j := i ; j < len(s); j++ {
            if set[int(s[j])] {break}
            set[int(s[j])] = true
            res = max(res, j-i+1)            
        }
    }
    return res
}