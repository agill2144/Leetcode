func lengthOfLongestSubstring(s string) int {
    res := 0
    if len(s) == 0 {return res}
    freq := map[byte]int{}
    left := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for left <= i && freq[s[i]] > 1 {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++            
        }
        res = max(res, i-left+1)
    }
    return res
}

// LOL brute force getting accepted
// func lengthOfLongestSubstring(s string) int {
//     res := 0
//     for i := 0; i < len(s); i++ {
//         set := make([]bool, 256)
//         for j := i ; j < len(s); j++ {
//             if set[int(s[j])] {break}
//             set[int(s[j])] = true
//             res = max(res, j-i+1)            
//         }
//     }
//     return res
// }