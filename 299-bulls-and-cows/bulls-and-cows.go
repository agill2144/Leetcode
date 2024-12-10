func getHint(secret string, guess string) string {
    if len(secret) != len(guess) {return ""}
    bulls := 0
    freq := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
            freq[secret[i]]++
        }
    }
    cows := 0
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {continue}
        if freq[guess[i]] != 0 {
            freq[guess[i]]--
            cows++
        }
    }
    return fmt.Sprintf("%vA%vB", bulls, cows)
}