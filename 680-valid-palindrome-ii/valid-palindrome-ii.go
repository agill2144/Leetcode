func validPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] == s[right] {
            left++
            right--
        } else {
            return isValid(s, left+1, right) || isValid(s, left, right-1)
        }
    }
    return true
}

func isValid(s string, left, right int) bool {
    for left <= right {
        if s[left] != s[right] {return false}
        left++
        right--
    }
    return true
}