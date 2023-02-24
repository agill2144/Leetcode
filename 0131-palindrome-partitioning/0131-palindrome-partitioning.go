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
        for i := start; i < len(s); i++ {
            subStr := string(s[start:i+1])
            if isPalindrome(subStr) {
                path = append(path, subStr)
                dfs(i+1, path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, []string{})
    return out
}

func isPalindrome(subStr string) bool {
    if len(subStr) == 1 {return true}
    left := 0
    right := len(subStr)-1
    for left < right {
        if subStr[left] != subStr[right] {
            return false
        }
        left++
        right--
    }
    return true
}