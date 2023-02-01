func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {return -1}
    if len(needle) == 0 {return -1}
    windowStr := ""
    left := 0
    for right := 0; right < len(haystack); right++ {
        windowStr += string(haystack[right])
        if right-left+1 == len(needle) {
            if windowStr[left:] == needle {
                return left
            }
            left++
        }
    }
    return -1
}