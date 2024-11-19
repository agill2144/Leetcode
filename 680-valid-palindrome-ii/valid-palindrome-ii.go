func validPalindrome(s string) bool {
    if len(s) <= 2 {return true}
    left := 0
    right := len(s)-1
    for left < right {
        if s[left] != s[right] {
            return isValid(left+1, right, s) || isValid(left, right-1, s)
        }
        left++;right--
    }
    return true
}

// xyabcbay
func isValid(left, right int, s string) bool {
    for left <= right {
        if s[left] != s[right] {return false}
        left++
        right--
    }
    return true
}