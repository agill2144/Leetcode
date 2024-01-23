func maxLength(arr []string) int {
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
    
    maxSize := 0
    var dfs func(start int, pathSize int, chars []bool)
    dfs = func(start int, pathSize int, chars []bool) {
        // base
        if pathSize > maxSize {maxSize = pathSize} 
        
        // logic
        for i := start; i < len(wordsWithUniqChars); i++ {
            word := wordsWithUniqChars[i]
            canUse := true
            for j := 0; j < len(word); j++ {
                if ok := chars[int(word[j]-'a')]; ok {
                    canUse = false
                    break
                }
            }
            
            if canUse {
                for j := 0; j < len(word); j++ {chars[int(word[j]-'a')] = true}
                dfs(i+1, pathSize+len(word), chars)
                for j := 0; j < len(word); j++ {chars[int(word[j]-'a')] = false}                
            } else {
                dfs(i+1, pathSize, chars)
            }
        }
    }
    dfs(0,0, make([]bool, 26))
    return maxSize
}