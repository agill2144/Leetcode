func firstPalindrome(words []string) string {
    for i := 0; i < len(words); i++ {
        if isPalindrome(words[i]) {return words[i]}
    }
    return ""   
}

func isPalindrome(s string)bool {
    left := 0
    right := len(s)-1
    for left < right {
        if s[left] != s[right]{return false}
        left++
        right--
    }
    return true
}