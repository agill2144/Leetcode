func strStr(haystack string, needle string) int {
    lps := findLPS(needle)
    i := 0
    j := 0
    for i < len(haystack) {
        if haystack[i] == needle[j] {
            j++
            i++
            if j == len(needle) {
                return i-len(needle)
            }
        } else if j > 0 {
            j = lps[j-1]
        } else {
            // j has reched 0th idx,
            // start matching from start of needle again
            i++
        }
    }
    return -1
}

func findLPS(needle string) []int {
    out := make([]int, len(needle))
    i := 1
    j := 0
    for i < len(needle) {
        if needle[i] == needle[j] {
            j++
            out[i] = j
            i++
        } else if j > 0 {
            j = out[j-1]
        } else if j == 0 {
            i++
        }
    }
    return out
}

/*
    approach: brute force
    - form needle size substring at each char of haystack
    
    time = o(h) * o(n)
    space = o(n) ; when forming needle size substr
*/
// func strStr(haystack string, needle string) int {
//     for i := 0; i <= len(haystack)-len(needle); i++ {
//         subStr := haystack[i:i+len(needle)]
//         if subStr == needle {return i}
//     }
//     return -1
// }