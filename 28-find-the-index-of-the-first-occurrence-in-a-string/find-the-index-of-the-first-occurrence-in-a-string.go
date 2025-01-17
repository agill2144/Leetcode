func strStr(s string, pattern string) int {
    if len(pattern) > len(s) {return -1}
    lps := calcLPS(pattern)
    i := 0
    j := 0
    for i < len(s) {
        if s[i] == pattern[j] {
            i++
            j++
            if j == len(pattern) {return i-j}
        } else {
            if j-1 < 0 {i++; continue}
            j = lps[j-1]
        }
    }
    return -1
}

func calcLPS(s string) []int {
    n := len(s)
    lps := make([]int, n)
    length := 0
    i := 1
    for i < n {
        if s[i] == s[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length-1 >= 0 {
                length = lps[length-1]
            } else {
                i++
            }
        }
    }
    return lps
}