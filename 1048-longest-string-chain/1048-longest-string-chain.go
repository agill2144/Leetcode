func longestStrChain(words []string) int {
    // sort.Slice(words, func(i, j int)bool{
    //     return len(words[i]) < len(words[j])
    // })

    // we can also bucket sort this
    // each word at max is of len 16
    // we map each len to idx which will contain 
    // a list words of that size
    bucket := make([][]string, 17)
    dpMap := map[string]int{}
    for i := 0; i < len(words); i++ {
        word := words[i]
        if bucket[len(word)] == nil {
            bucket[len(word)] = []string{}
        }
        bucket[len(word)] = append(bucket[len(word)], word)
        dpMap[word] = 1
    }
    
   
    
    longestChain := 1
    // why start from 1, because for words in bucket[0],
    // their best answer WILL ALWAYS BE size 1
    // because there is nothing to look back for words at 0th position
    // to try and form other smaller substrings
    for i := 1; i < len(bucket); i++ {
        for _, word := range bucket[i] {
            if len(word) == 1 {continue}
            for j := 0; j < len(word); j++ {
                subStr := word[0:j] + word[j+1:]
                if count, ok := dpMap[subStr]; ok {
                    dpMap[word] = max(count+1, dpMap[word])
                    longestChain = max(dpMap[word], longestChain)
                }
            }
        }
    }
    return longestChain
}

func max(x, y int) int {
    if x > y {return x}
    return y
}