func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {return ""}
    if len(strs) == 1 {return strs[0]}
    
    for i := 0; i < len(strs[0]); i++ {
        for j := 1; j < len(strs) ; j++ {
            if i == len(strs[j]) || strs[j][i] != strs[0][i] {return strs[0][:i]}
        }
    }
    return strs[0]
}
/*
    approach: sort + 2 ptrs
    - sorting groups common prefixes
    - then take 2 ptrs to compare first str and last str
        - compare char by char
        - we have to compare 1st and last because
        - we want to find the common prefix that is present in ALL strings
    n = len(strs)
    k = avg len of each word
    time = o(nlogn) + o(k)
    space = o(1)
*/
// func longestCommonPrefix(strs []string) string {
//     n := len(strs)
//     if n <= 1 {
//         if n == 1 {return strs[0]}
//         return ""
//     }
//     sort.Strings(strs)
//     left := 0
//     right := 0
//     for left < len(strs[0]) && right < len(strs[n-1]) {
//         if strs[0][left] != strs[len(strs)-1][right] {
//             return strs[0][:left]
//         }
//         left++
//         right++
//     }
//     return strs[0]
// }
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