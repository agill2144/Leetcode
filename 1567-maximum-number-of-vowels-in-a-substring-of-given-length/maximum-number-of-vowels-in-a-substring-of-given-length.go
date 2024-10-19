func maxVowels(s string, k int) int {
    count := 0
    vCount := 0
    isV := func(char byte) bool {
        return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
    }
    left := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        if isV(char) {
            vCount++
        }

        if i-left+1 == k {
            count = max(count, vCount)
            if isV(s[left]) {vCount--}
            left++            
        }
    }
    return count
}