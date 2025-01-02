func vowelStrings(words []string, queries [][]int) []int {
    n := len(words)
    vCount := make([]int, n)
    for i := 0; i < n; i++ {
        prev := 0
        if i-1 >= 0 {prev = vCount[i-1]}
        vCount[i] = prev
        if meetsVowel(words[i]) {vCount[i]++}
    }
    out := make([]int, len(queries))
    for i := 0; i < len(queries); i++ {
        left, right := queries[i][0], queries[i][1]
        count := vCount[right]
        leftCount := 0
        if left-1 >= 0 {leftCount = vCount[left-1]}
        count -= leftCount
        out[i] = count
    }
    return out
}

func meetsVowel(word string) bool {
    v := map[byte]bool{'a':true,'e':true,'i':true,'o':true,'u':true}
    return v[word[0]] && v[word[len(word)-1]]
}