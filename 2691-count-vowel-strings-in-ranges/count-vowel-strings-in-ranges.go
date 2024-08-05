func vowelStrings(words []string, queries [][]int) []int {
    vowels := make([]bool, len(words))
    prefixCounts := make([]int, len(words))
    for i := 0; i < len(words); i++ {
        vowels[i] = isVowelWord(words[i])
        prevCount := 0
        if i - 1 >= 0 {prevCount = prefixCounts[i-1]}
        prefixCounts[i] = prevCount
        if vowels[i] {prefixCounts[i]++}
    }
    out := []int{}
    for i := 0; i < len(queries); i++ {
        start := queries[i][0]
        end := queries[i][1]
        x := prefixCounts[end]
        y := 0
        if start-1 >= 0 {y = prefixCounts[start-1]}
        out = append(out, x-y)
    }
    return out
}

func isVowelWord(word string) bool {
    vowels := map[byte]bool{
        'a': true, 'e': true, 'i':true, 'o':true, 'u':true,
    }
    return vowels[word[0]] && vowels[word[len(word)-1]]
}