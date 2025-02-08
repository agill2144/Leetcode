func maxLength(arr []string) int {
    if len(arr) == 0 {return 0}
    words := []string{}
    for i := 0; i < len(arr); i++ {
        set := map[byte]bool{}
        hasDupeChars := false
        word := arr[i]
        for k := 0; k < len(arr[i]); k++ {
            if set[word[k]] {hasDupeChars = true; break}
            set[word[k]] = true
        }
        if !hasDupeChars {words = append(words, word)}
    }
    res := 0
    if len(words) == 0 {return 0}
    var dfs func(start int, path []bool, currSize int)
    dfs = func(start int, path []bool, currSize int) {
        // base
        res = max(res, currSize)
        
        // logic
        for i := start; i < len(words); i++ {
            word := words[i]
            canBeUsed := true
            for k := 0; k < len(word); k++ {
                charIdx := int(word[k]-'a')
                if path[charIdx] {canBeUsed = false; break}
            }
            if canBeUsed {
                for k := 0; k < len(word); k++ { path[int(word[k]-'a')] = true}
                dfs(i+1, path, currSize + len(word))
                for k := 0; k < len(word); k++ { path[int(word[k]-'a')] = false}
            }
        }
    }
    dfs(0, make([]bool, 26), 0)
    return res
}