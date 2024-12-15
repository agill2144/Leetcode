/*
    approach: brute force
    - create all possible substr from strs[0]; starting from size 1, then 2, then 3
    - and check if this substr exists from idx 1 -> n-1 in rest of the word
*/
func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n <= 1 {
        if n == 0 {return ""}
        return strs[0]
    }
    word := strs[0]
    // loop over reserved word
    // to form substrs
    ans := ""
    for i := 0; i < len(word); i++ {
        prefix := word[:i+1]
        count := 0
        for j := 1; j < len(strs); j++ {
            if strings.HasPrefix(strs[j], prefix) {
                count++
                if count == n-1 {
                    ans = prefix
                    break
                }
            }
        }
    }
    return ans
}