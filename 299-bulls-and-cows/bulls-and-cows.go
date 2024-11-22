func getHint(secret string, guess string) string {
    bulls := 0
    freq := map[byte]int{}
    for i := 0; i < len(guess); i++ {
        freq[secret[i]]++
        if secret[i] == guess[i] {bulls++; freq[secret[i]]--}
    }
    cows := 0
    for i := 0; i < len(guess); i++ {
        if secret[i] == guess[i] {continue}
        count := freq[guess[i]]
        if count > 0 {
            cows++; freq[guess[i]]--
        }
    }
    return fmt.Sprintf("%vA%vB", bulls, cows)
}
