func getHint(secret string, guess string) string {
    gFreq := make([]int, 10)
    sFreq := make([]int, 10)
    bulls := 0
    cows := 0
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
            if sFreq[int(guess[i]-'0')] > 0 {
                sFreq[int(guess[i]-'0')]--
                cows++
            } else {
                gFreq[int(guess[i])-'0']++
            }
            
            if gFreq[int(secret[i]-'0')] > 0 {
                gFreq[int(secret[i]-'0')]--
                cows++
            } else {
                sFreq[int(secret[i])-'0']++                
            }
        }
    }
    return fmt.Sprintf("%vA%vB",bulls,cows)
}