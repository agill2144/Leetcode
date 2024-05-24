func maxScoreWords(words []string, letters []byte, score []int) int {
    maxScore := 0
        
    lettersFreq := map[byte]int{}
    for i := 0; i < len(letters); i++ {
        lettersFreq[letters[i]]++
    }

    var dfs func(start int, sum int) 
    dfs = func(start int, sum int) {
        // base
        maxScore = max(maxScore, sum)

        // logic
        // n words 
        // o(n)
        for i := start; i < len(words); i++ {
            word := words[i]
            wordSum := 0

            // k avg size of each word
            // o(k + 26+26)
            wordFreq := map[byte]int{}
            for j := 0;  j < len(word); j++ {
                char := word[j]
                wordFreq[char]++
                wordSum += score[int(char-'a')]
            }

            isPossible := true
            // o(26)
            for char,countNeeded := range wordFreq {
                countAvail, charExists := lettersFreq[char]
                if !charExists || countAvail < countNeeded {
                    isPossible = false
                    break
                }
            }

            if isPossible {
                // o(26)
                for char, val := range wordFreq {lettersFreq[char]-=val}
                newSum := wordSum + sum
                dfs(i+1, newSum)
                // o(26)
                for char, val := range wordFreq {lettersFreq[char]+=val}
            }
        }
    }
    dfs(0,0)
    return maxScore
}