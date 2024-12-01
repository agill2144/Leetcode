func isPalindrome(s string) bool {
    left := 0
    right := len(s)-1
    for left < right {
        for left < right && !isAlphaNumeric(s[left]) {left++}
        if left == right {break}
        for left < right && !isAlphaNumeric(s[right]) {right--}
        if left == right {break}
        if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {return false}
        left++
        right--
    }
    return true
}

func isAlphaNumeric(char byte) bool {
    return (char >= '0' && char <= '9') || 
            (char >= 'a' && char <= 'z') || 
            (char >= 'A' && char <= 'Z')
}


