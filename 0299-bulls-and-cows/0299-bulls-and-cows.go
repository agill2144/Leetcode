func getHint(secret string, guess string) string {
    digitCount := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        digitCount[secret[i]]++
    }
    bulls := 0
    for i := 0; i < len(guess); i++ {
        if secret[i] == guess[i] {
            bulls++
            digitCount[guess[i]]--
        }
    }
    cows := 0
    for i := 0; i < len(guess); i++ {
        count := digitCount[guess[i]]
        if guess[i] != secret[i] && count > 0 {
            cows++
            digitCount[guess[i]]--
        }
    }
    return fmt.Sprintf("%vA%vB", bulls, cows)
}