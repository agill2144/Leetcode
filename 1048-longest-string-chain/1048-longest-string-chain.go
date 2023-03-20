func longestStrChain(words []string) int {
    sort.Slice(words, func(i, j int)bool{
        return len(words[i]) < len(words[j])
    })

    dpMap := map[string]int{}
    for i := 0; i < len(words); i++ {dpMap[words[i]] = 1}
    longestChain := 1
    for i := 0; i < len(words); i++ {
        word := words[i]
        
        if len(word) == 1 {continue}
        
        for j := 0; j < len(word); j++ {
            subStr := word[0:j] + word[j+1:]
            if count, ok := dpMap[subStr]; ok {
                dpMap[word] = max(count+1, dpMap[word])
                longestChain = max(dpMap[word], longestChain)
            }
        }
        
        
    }
    return longestChain
}

func max(x, y int) int {
    if x > y {return x}
    return y
}