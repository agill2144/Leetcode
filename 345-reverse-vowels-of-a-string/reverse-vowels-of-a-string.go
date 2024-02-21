func reverseVowels(s string) string {
    vowels := map[string]struct{}{"a":{},"e":{},"i":{},"o":{},"u":{}}
    isVowel := func(char string) bool {
        _, ok := vowels[strings.ToLower(char)]
        return ok
    }
    sList := strings.Split(s, "")
    left := 0
    right := len(s)-1
    for left < right {

        for left < right && !isVowel(sList[left]) {
            left++
        }
        for left < right && !isVowel(sList[right]) {
            right--
        }

        sList[left],sList[right] = sList[right],sList[left]
        left++; right--
    }
    return strings.Join(sList, "")
}