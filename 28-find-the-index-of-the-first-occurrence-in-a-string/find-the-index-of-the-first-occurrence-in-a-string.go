func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {return -1}
    lps := calcLPS(needle)
    i, j := 0, 0
    for i < len(haystack) && j < len(needle) {
        if haystack[i] == needle[j] {
            i++
            j++
        } else {
            if j > 0 {
                j = lps[j-1]
            } else {
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
    j := 0 // prefix window ptr ( prefix win size & incoming char in prefix win )
    i := 1 // incoming char in suffix win
    for i < n {
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