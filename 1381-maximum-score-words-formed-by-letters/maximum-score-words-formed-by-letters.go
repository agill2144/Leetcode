func maxScoreWords(words []string, letters []byte, score []int) int {
    maxScore := 0
        
    lettersFreq := map[byte]int{}
    for i := 0; i < len(letters); i++ {
        lettersFreq[letters[i]]++
    }

    var dfs func(start int, sum int, freq map[byte]int) 
    dfs = func(start int, sum int, freq map[byte]int) {
        // base
        maxScore = max(maxScore, sum)

        // logic
        for i := start; i < len(words); i++ {
            word := words[i]
            wordSum := 0
            
            wordFreq := map[byte]int{}
            for j := 0;  j < len(word); j++ {
                char := word[j]
                wordFreq[char]++
                wordSum += score[int(char-'a')]
            }

            isPossible := true
            for char,countNeeded := range wordFreq {
                countAvail, charExists := freq[char]
                if !charExists || countAvail < countNeeded {
                    isPossible = false
                    break
                }
            }

            if isPossible {
                // fmt.Println(word,"------------------")
                // fmt.Println(freq)
                for char, val := range wordFreq {freq[char]-=val}
                newSum := wordSum + sum
                // fmt.Println(freq, newSum)
                dfs(i+1, newSum, freq)
                for char, val := range wordFreq {freq[char]+=val}
                // fmt.Println(freq)
                // fmt.Println("--#############################")
            }
        }
    }
    dfs(0,0, lettersFreq)
    return maxScore
}