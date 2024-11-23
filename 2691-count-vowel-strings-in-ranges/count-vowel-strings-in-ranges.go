func vowelStrings(words []string, queries [][]int) []int {
    prefixCounts := make([]int, len(words))
    for i := 0; i < len(words); i++ {
        prev := 0
        if i-1 >= 0 {prev = prefixCounts[i-1]}
        if meetsVowelCondition(words[i]) {            
            prefixCounts[i] = prev+1
        } else {
            prefixCounts[i] = prev
        }
    }
    out := []int{}
    for i := 0; i < len(queries); i++ {
        start, end := queries[i][0], queries[i][1]
        total := prefixCounts[end]
        if start - 1 >= 0 {total -= prefixCounts[start-1]}
        out = append(out, total)
    }
    return out
}

func meetsVowelCondition(word string) bool {
    // words that start and end with a vowel
    vs := map[byte]bool{'a':true,'e':true,'i':true,'o':true,'u':true}
    return vs[word[0]] && vs[word[len(word)-1]]
}