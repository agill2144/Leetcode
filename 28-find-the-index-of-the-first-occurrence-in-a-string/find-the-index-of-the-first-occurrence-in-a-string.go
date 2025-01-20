/*
    approach: brute force
    - try evaluating each position in haystack string whether its the first index or not
    time = o
    approach: kmp
    - instead
*/
func strStr(haystack string, needle string) int {
    for i := 0; i < len(haystack); i++ {
        h := i
        n := 0
        for h < len(haystack) && n < len(needle) {
            if haystack[h] != needle[n] {break}
            h++
            n++
        }
        if n == len(needle) {return h-len(needle)}
    }
    return -1
}