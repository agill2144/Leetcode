// The most magical fucking algo : KMP
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
        } else if haystack[i] != needle[j] && j > 0 {
            /*
                When you encounter a mismatch at haystack[i] and needle[j],
                You adjust j based on the LPS (Longest Prefix Suffix) array. 
                The LPS array helps you determine how much of the pattern
                can be reused when a mismatch occurs.
            */
            j = lps[j-1]
        } else if haystack[i] != needle[j] && j == 0 {
            // j is already at 0th idx, cant move that back anymore
            i++
        }
    }
    return -1
}

func findLPS(needle string) []int {
    out := make([]int, len(needle))
    j := 0
    i := 1
    for i < len(needle) {
        if needle[i] == needle[j] {
            j++
            out[i] = j
            i++
        } else if needle[j] != needle[i] && j > 0 {
            j = out[j-1]
        } else if needle[j] != needle[i] && j == 0 {
            i++
        }
    }
    return out
}