func getHint(secret string, guess string) string {
    bulls := 0
    cows := 0
    sFreqMap := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        sFreqMap[secret[i]]++
    }
    
    for i := 0; i < len(guess); i++ {
        char := guess[i]
        if count, ok := sFreqMap[char]; ok {
            if char == secret[i] {
                bulls++
                if val := sFreqMap[char]; val <= 0 {cows--} 
            } else if count > 0 { cows++ }
            sFreqMap[char]--

        }
    }
    
    return fmt.Sprintf("%vA%vB", bulls,cows)
}







