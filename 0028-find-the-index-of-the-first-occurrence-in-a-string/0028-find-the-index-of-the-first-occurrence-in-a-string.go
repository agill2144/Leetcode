// The most magical fucking algo : KMP
func strStr(haystack string, needle string) int {
    lps := lps(needle)
    i := 0 // haystack ptr
    j := 0 // needle ptr
    for i < len(haystack) {
        if haystack[i] == needle[j] {
            j++
            i++
            if j == len(needle) {
                return i - len(needle)
            }
        } else if haystack[i] != needle[j] && j > 0 {
            j = lps[j-1]
        } else if haystack[i] != needle[j] && j == 0{
            i++ // j is already at 0th idx, cant move that back anymore
        }
    }
    return -1
}

func lps(needle string) []int {
    out := make([]int, len(needle))
    i := 1 // suffix window right ptr
    j := 0 // prefix window right ptr
    for i < len(needle) {
        if needle[i] == needle[j] {
            j++
            out[i] = j
            i++
        } else if needle[i] != needle[j] && j > 0 {
            j = out[j-1]
        } else if needle[i] != needle[j] && j == 0 {
            i++
        }
    }
    return out
}