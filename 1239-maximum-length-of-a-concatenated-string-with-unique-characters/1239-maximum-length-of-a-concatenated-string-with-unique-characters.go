func maxLength(arr []string) int {
    
    // 1. filter out words with duplicate chars
    wordsWithUniqChars := []string{}
    for i := 0; i < len(arr); i++ {
        word := arr[i]
        chars := make([]bool, 26)
        isUniq := true
        for j := 0; j < len(word); j++ {
            idx := int(word[j]-'a')
            ok := chars[idx]
            if ok {
                isUniq = false
                break
            }
            chars[idx] = true
        }
        
        if isUniq {
            wordsWithUniqChars = append(wordsWithUniqChars, word)
        }
    }
    
    // 2. use dfs + backtracking, to form all possible subseq 
    maxSize := 0
    var dfs func(start int, pathSize int, chars []bool)
    dfs = func(start int, pathSize int, chars []bool) {
        // base
        if pathSize > maxSize {maxSize = pathSize} 
        
        // logic
        for i := start; i < len(wordsWithUniqChars); i++ {
            word := wordsWithUniqChars[i]
            
            // its possible that we cannot use this word as some of the chars
            // from this word may already be in our path
            // check if this word can be used 
            canUse := true
            for j := 0; j < len(word); j++ {
                if ok := chars[int(word[j]-'a')]; ok {
                    canUse = false
                    break
                }
            }
            
            // if we can use it, add its chars to our path
            // recurse with updated pathSize
            if canUse {
                // action
                for j := 0; j < len(word); j++ {chars[int(word[j]-'a')] = true}
                // recurse
                dfs(i+1, pathSize+len(word), chars)
                // backtrack
                for j := 0; j < len(word); j++ {chars[int(word[j]-'a')] = false}
            }
            // if we cant use this word, try the next ith word, continue for loop 
        }
    }
    dfs(0,0, make([]bool, 26))
    return maxSize
}