// time = o(w)
// space = o(1)
func numberOfSpecialChars(word string) int {
    // instead of using a set to only store 
    // uniq pairs
    // we can mark each lower and upper case chars used 
    lowers := make([]bool, 26)
    uppers := make([]bool, 26)
    seen := map[byte]struct{}{}
    count := 0
    for i := 0; i < len(word); i++ {
        char := word[i]
        if unicode.IsUpper(rune(char)) {
            // look for lower
            lower := char+32
            upper := char
            _, ok := seen[lower]
            if ok && !lowers[int(lower-'a')] && !uppers[int(upper-'A')] {
                lowers[int(lower-'a')] = true
                uppers[int(upper-'A')] = true
                count++
            }
        } else if unicode.IsLower(rune(char)) {
            // look for upper
            lower := char
            upper := char-32
            _, ok := seen[upper]
            if ok && !lowers[int(lower-'a')] && !uppers[int(upper-'A')] {
                lowers[int(lower-'a')] = true
                uppers[int(upper-'A')] = true
                count++
            }
        }
        seen[char] = struct{}{}
    }
    return count
}