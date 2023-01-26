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
            subStr := s[start:i+1]
            if !isPalindrome(subStr) {continue}
            path = append(path, subStr)
            dfs(i+1, path)
            path = path[:len(path)-1]
        }
    }
    dfs(0, nil)
    return out
    
}

func isPalindrome(s string) bool {
    left := 0
    right := len(s)-1
    for left < right {
        if s[left] == s[right] {
            left++
            right--
        } else {
            return false
        }
    }
    return true
}