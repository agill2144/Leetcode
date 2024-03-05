func minimumLength(s string) int {
    left := 0
    right := len(s)-1
    for left < right {
        if s[left] == s[right] {
            left++
            right--
            for left <= right && s[left] == s[left-1] {left++}
            for left <= right && s[right] == s[right+1] {right--}
        } else {
            break
        }
    }
    return right-left+1
}