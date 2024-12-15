func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n <= 1 {
        if n == 1 {return strs[0]}
        return ""
    }
    sort.Strings(strs)
    left := 0
    right := 0
    for left < len(strs[0]) && right < len(strs[n-1]) {
        if strs[0][left] != strs[len(strs)-1][right] {
            return strs[0][:left]
        }
        left++
        right++
    }
    return strs[0]
}
/*
    approach: brute force
    num of words = n
    avg size of each word = k
    tc = o(k * n*k)
    sc = o(1)
*/
// func longestCommonPrefix(strs []string) string {
//     if len(strs) <= 1 {
//         if len(strs) == 1 {return strs[0]}
//         return ""
//     }
//     word := strs[0]
//     for i := 0; i < len(word); i++ {
//         prefix := word[0:len(word)-i]
//         count := 0
//         for j := 1; j < len(strs); j++ {
//             if strings.HasPrefix(strs[j], prefix) {
//                 count++
//                 if count == len(strs)-1 {return prefix}
//             }
//         }
//     }
//     return ""
// }