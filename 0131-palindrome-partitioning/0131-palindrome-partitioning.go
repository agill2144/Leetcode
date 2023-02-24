func partition(s string) [][]string {
    out := [][]string{}
    var dfs func(start int, path []string)
    dfs = func(start int, path []string) {
        // base
        if start == len(s) {
            newL := make([]string, len(path))
            copy(newL, path)
            out = append(out, newL)
            return
        }
        
        // logic
        for i := start; i<len(s); i++ {
            subStr := string(s[start:i+1])
            if isPalindrome(subStr) {
                // action
                path = append(path, subStr)
                // recurse
                dfs(i+1, path)
                // backtrack
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, []string{})
    return out
}

func isPalindrome(s string) bool {
    if len(s) <= 1 {return true}
    left := 0; right := len(s)-1
    for left < right {
        if s[left] != s[right] {return false}
        left++
        right--
    }
    return true
}