func isPalindrome(s string) bool {
    var dfs func(left, right int)bool
    dfs = func(left, right int)bool {
        // base
        if left > right {return true}

        // logic
        leftChar := s[left]
        rightChar := s[right]
        if !isAlpha(leftChar) {
            return dfs(left+1, right)
        }
        if !isAlpha(rightChar) {
            return dfs(left, right-1)
        }
        leftStr := strings.ToLower(string(leftChar))
        rightStr := strings.ToLower(string(rightChar))
        fmt.Println(leftStr, rightStr)
        return leftStr == rightStr && dfs(left+1, right-1)
    }
    return dfs(0, len(s)-1)
}

func isAlpha(c byte)bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}