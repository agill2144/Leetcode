func isPalindrome(s string) bool {
    l := 0
    r := len(s)-1
    for l < r {
        for l < r && !isAlphaNumeric(s[l]) {l++}
        if l == r {break}
        
        for l < r && !isAlphaNumeric(s[r]) {r--}
        if l == r {break}

        if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {return false}
        l++
        r--
    }
    return true
}

func isAlphaNumeric(char byte)bool {
    return (char >= '0' && char <= '9') || 
            (char >= 'a' && char <= 'z') ||
            (char >= 'A' && char <= 'Z')
}