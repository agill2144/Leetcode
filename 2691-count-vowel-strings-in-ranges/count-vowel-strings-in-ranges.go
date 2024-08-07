func vowelStrings(words []string, queries [][]int) []int {
    prefixCounts := make([]int, len(words))
    for i := 0; i < len(words); i++ {
        if isVowel(words[i]) {
            count := 1
            if i-1 >= 0 {count += prefixCounts[i-1]}
            prefixCounts[i] = count
        } else {
            if i-1 >= 0 {prefixCounts[i] = prefixCounts[i-1]}            
        }
    }
    out := make([]int, len(queries))
    for i := 0; i < len(queries); i++ {
        start := queries[i][0]
        end := queries[i][1]
        x := prefixCounts[end]
        y := 0
        if start - 1 >= 0 { y = prefixCounts[start-1] }
        out[i] = x-y
    }
    return out
}

func isVowel(word string) bool {
    char1 := word[0]
    char2 := word[len(word)-1]
    vowels := map[byte]bool{'a':true,'e':true,'i':true,'o':true,'u':true}
    return vowels[char1] && vowels[char2]
}