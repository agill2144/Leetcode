/*
    approach: vertical scanning
    - instead of creating $evalPrefix
    - and then comparing with rest of words,
    - we can just compare whether ith idx char in reserved word 
        matches ith idx in rest of the strings
    - so this still means, reserver a word first
        - strs[0]
    - loop over this reserved word, char by char - ith idx
    - then loop over remaining strs from index 1
    - and check if ith idx char == reserved char ith idx
    - if it does not equal, ith idx was tracking the common prefix so far
        - so return reservedWord[:i]
        - substr up-to but not including i
    S = total number of character in all strs
    tc = o(S)
    sc = o(1)
*/
func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n <= 1 {
        if n == 1 {return strs[0]}
        return ""
    }
    reserved := strs[0]
    for i := 0; i < len(reserved); i++ {
        for j := 1; j < len(strs); j++ {
            word := strs[j]
            if i == len(word) || word[i] != reserved[i] {return reserved[:i]}
        }
    }
    return reserved
}
/*
    approach: brute force
    - create all possible substr from strs[0]; starting from size 1, then 2, then 3
        - $evalPrefix
    - and check if this substr as a prefix in rest of the strings
        - that is, start from idx 1
        - check if strs[j] has a prefix of $evalPrefix
        - if it does, increment a counter and if count reaches n-1
            - then is a good answer, save it
            - and try a longer $evalPrefix
            - because we want longest common
    n = len(strs)
    k = avg size of each strs[i]
    tc = o(k * n*k) = o(n*k^2)
    sc = o(1)
*/
// func longestCommonPrefix(strs []string) string {
//     n := len(strs)
//     if n <= 1 {
//         if n == 0 {return ""}
//         return strs[0]
//     }
//     word := strs[0]
//     // loop over reserved word
//     // to form substrs
//     ans := ""
//     for i := 0; i < len(word); i++ {
//         prefix := word[:i+1]
//         count := 0
//         for j := 1; j < len(strs); j++ {
//             if strings.HasPrefix(strs[j], prefix) {
//                 count++
//                 if count == n-1 {
//                     ans = prefix
//                     break
//                 }
//             }
//         }
//     }
//     return ans
// }