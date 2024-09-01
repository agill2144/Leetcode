func vowelStrings(words []string, queries [][]int) []int {
    n := len(words)
    prefixCounts := make([]int, n)
    for i := 0; i < len(words); i++ {
        count := 0
        if i-1 >= 0 {count = prefixCounts[i-1]}
        if isVowel(words[i]) {
            count++
        }
        prefixCounts[i] = count
    }
    out := []int{}
    for i := 0; i < len(queries); i++ {
        left := queries[i][0]
        right := queries[i][1]
        fullCount := prefixCounts[right]
        leftCount := 0
        if left-1 >= 0 {leftCount = prefixCounts[left-1]}
        out = append(out, fullCount-leftCount)
    }        
    return out
}

func isVowel(word string) bool {
    vMap := map[byte]bool{
        'a':true,'e':true,'i':true,'o':true,'u':true,
    }
    return vMap[word[0]] && vMap[word[len(word)-1]]
}