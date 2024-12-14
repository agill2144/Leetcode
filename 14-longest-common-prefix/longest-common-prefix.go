func longestCommonPrefix(strs []string) string {
    if len(strs) <= 1 {
        if len(strs) == 1 {return strs[0]}
        return ""
    }
    word := strs[0]
    for i := 0; i < len(word); i++ {
        prefix := word[0:len(word)-i]
        count := 0
        for j := 1; j < len(strs); j++ {
            if strings.HasPrefix(strs[j], prefix) {
                count++
                if count == len(strs)-1 {return prefix}
            }
        }
    }
    return ""
}