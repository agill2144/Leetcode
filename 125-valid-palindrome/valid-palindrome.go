func isPalindrome(s string) bool {
    t := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if (s[i] >= '0' && s[i] <= '9') || 
            (s[i] >= 'a' && s[i] <= 'z') ||
            (s[i] >= 'A' && s[i] <= 'Z') {
                t.WriteString(strings.ToLower(string((s[i]))))
            }
    }
    tStr := t.String()
    l := 0
    r := len(tStr)-1
    for l < r {
        if tStr[l] != tStr[r] {return false}
        l++
        r--
    }
    return true
}