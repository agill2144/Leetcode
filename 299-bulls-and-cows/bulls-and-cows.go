func getHint(secret string, guess string) string {
    g := make([]int, 10)
    s := make([]int, 10)
    bulls := 0
    cows := 0
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
            if s[int(guess[i]-'0')] > 0 {
                s[int(guess[i]-'0')]--
                cows++
            } else {
                g[int(guess[i])-'0']++
            }
            
            if g[int(secret[i]-'0')] > 0 {
                g[int(secret[i]-'0')]--
                cows++
            } else {
                s[int(secret[i])-'0']++                
            }
        }
    }
    return fmt.Sprintf("%vA%vB",bulls,cows)
}