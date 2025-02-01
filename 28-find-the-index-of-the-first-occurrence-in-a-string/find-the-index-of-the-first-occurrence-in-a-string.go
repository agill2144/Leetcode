func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {return -1}
    lps := calcLPS(needle)
    i := 0 // haystack ptr
    j := 0 // needle ptr
    for i < len(haystack) && j < len(needle) {
        if haystack[i] == needle[j] {
            i++
            j++
        } else {
            if j-1 >= 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    if j == len(needle) {
        return i-len(needle)
    }
    return -1
}

func calcLPS(word string) []int {
    n := len(word)
    lps := make([]int, n)
    j := 0
    i := 1
    for i < len(word) {
        if word[i] == word[j] {
            j++
            lps[i] = j
            i++
        } else {
            if j-1 >= 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    return lps
}