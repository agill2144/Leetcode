func partition(s string) [][]string {
    out := [][]string{}
    var dfs func(start int, path []string)
    dfs = func(start int, path []string) {
        // base
        if start == len(s) {
            newP := make([]string, len(path))
            copy(newP, path)
            out = append(out, newP)
            return
        }

        // logic
        for i := start; i < len(s); i++ {
            subStr := s[start:i+1]
            if isPalindrome(subStr) {
                path = append(path, subStr)
                dfs(i+1, path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, nil)
    return out
}


func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for l < r {
        if s[l] != s[r] {return false}
        l++
        r--
    }
    return true
}