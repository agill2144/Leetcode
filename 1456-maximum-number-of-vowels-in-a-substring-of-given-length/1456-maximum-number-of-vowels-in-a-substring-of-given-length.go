func maxVowels(s string, k int) int {
    maxCount := 0
    vowels := map[byte]struct{}{'a':{},'e':{},'i':{},'o':{},'u':{}}
    count := 0
    left := 0
    for i := 0; i < len(s); i++ {
        if _, ok := vowels[s[i]]; ok {count++}
        if i-left+1 == k {
            if count > maxCount {maxCount = count}
            if _, ok := vowels[s[left]]; ok {count--}
            left++
        }
    }
    return maxCount
}