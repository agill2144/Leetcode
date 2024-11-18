func validPalindrome(s string) bool {
    if len(s) <= 2 {return true}
    left := 0
    right := len(s)-1
    for left < right {
        if s[left] == s[right] {
            left++; right--
        } else {
            return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
        }
    }
    return true
}

func isPalindrome(s string, left, right int) bool {
    if left > right {return false}
    for left < right {
        if s[left] != s[right] {return false}
        left++
        right--
    }
    return true
}