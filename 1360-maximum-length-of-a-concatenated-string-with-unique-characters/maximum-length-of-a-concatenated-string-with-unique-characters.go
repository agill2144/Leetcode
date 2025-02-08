func maxLength(arr []string) int {
    // only use words with uniq chars
    // ["abc", "def", "ghhi"]
    // we can never make anything with "ghhi" - because it has duplicate h's
    // therefore, 1st find words with uniq chars
    uniqWords := []string{}
    for i := 0; i < len(arr); i++ {
        word := arr[i]
        seen := make([]bool, 26)
        isUniq := true
        for j := 0; j < len(word); j++ {
            if seen[int(word[j]-'a')] {
                isUniq = false
                break
            }
            seen[int(word[j]-'a')] = true
        }
        if isUniq {
            uniqWords = append(uniqWords, word)
        }
    }
    
    // now its just dfs with backtracking
    // we will be backtracking "seen" bool vars to keep track of what all characters we have seen so far
    // this way when we are evaluating a word to use, we will go thru all chars of that word first to check
    // whether we can use this word or not ( all chars of this word have not been seen yet )
    longest := 0
    var dfs func(start int, path int, seen []bool)
    dfs = func(start int, path int, seen []bool) {
        // base
        longest = max(longest, path)
        
        // logic
        for i := start; i < len(uniqWords); i++ {
            word := uniqWords[i]

            canUse := true
            for j := 0; j < len(word); j++ {
                idx := int(word[j]-'a')
                if seen[idx] {
                    canUse = false
                    break
                }
            }
            
            if canUse {
                // mark each char seen
                for j := 0; j < len(word); j++ { seen[int(word[j]-'a')] = true  }
                dfs(i+1, path+len(word), seen)
                // mark each char un-seen
                for j := 0; j < len(word); j++ { seen[int(word[j]-'a')] = false  }
                
            }
            
        }
    }
    dfs(0, 0, make([]bool, 26))
    return longest
}