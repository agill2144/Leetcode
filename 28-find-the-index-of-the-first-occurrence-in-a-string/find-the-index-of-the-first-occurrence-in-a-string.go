func strStr(haystack string, needle string) int {
    if len(haystack) < len(needle) {return -1}
    lps := calcLPS(needle)
    i, j := 0, 0
    for i < len(haystack) && j < len(needle) {
        if haystack[i] == needle[j] {
            i++
            j++
        } else {
            // what is a longest COMMON suffix before this char
            // that is also a prefix, that we can skip 
            // checking because they are equal
            if j != 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    if j == len(needle) {return i-len(needle)}
    return -1
}

func calcLPS(word string) []int{
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
            if j == 0 {
                i++
            } else {
                j = lps[j-1]
            }
        }
    }
    return lps
}