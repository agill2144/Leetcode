func getHint(secret string, guess string) string {
    bulls := 0
    cows := 0
    freq := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {bulls++; continue}
        freq[secret[i]]++
    }
    for i := 0; i < len(guess); i++ {
        if secret[i] == guess[i] {continue}
        if freq[guess[i]] > 0 {
            cows++
            freq[guess[i]]--
        }
    }
    return fmt.Sprintf("%vA%vB",bulls, cows)
}
