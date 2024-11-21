func validPalindrome(s string) bool {
    left := 0
    right := len(s)-1
    for left < right {
        if s[left] != s[right] {
            return isValid(left+1, right, s) || isValid(left, right-1, s)
        }
        left++
        right--
    }
    return true
}

func isValid(left, right int, s string) bool {
    if right-left+1 <= 1 {return true}
    for left <= right {
        if s[left] != s[right] {return false}
        left++
        right--
    }
    return true
}