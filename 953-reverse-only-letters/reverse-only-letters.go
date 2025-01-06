func reverseOnlyLetters(s string) string {
    sList := strings.Split(s, "")
    left := 0
    right := len(sList)-1
    for left < right {
        if !isEnglish(sList[left]) {left++; continue}
        if !isEnglish(sList[right]) {right--; continue}
        sList[left], sList[right] = sList[right], sList[left]
        left++
        right--
    }
    return strings.Join(sList,"")
}

func isEnglish(c string) bool {
    return (c >= "a" && c <=  "z") || (c >= "A" && c <= "Z")
}
