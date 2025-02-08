func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) { return -1 }
    lps := calcLPS(needle)
    i := 0
    j := 0
    for i < len(haystack) && j < len(needle) {
        if haystack[i] == needle[j] {
            i++
            j++
        } else {
            if j > 0 {
                j = lps[j-1]
            } else {
                // j is at 0, start a fresh match
                i++
            }
        }
    }
    if j == len(needle) {return i-len(needle)}
    return -1
}

func calcLPS(word string) []int {
    n := len(word)
    lps := make([]int, n)
    i := 1
    j := 0
    for i < len(word) {
        if word[i] == word[j] {
            j++
            lps[i] = j
            i++
        } else {
            if j > 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    return lps
}