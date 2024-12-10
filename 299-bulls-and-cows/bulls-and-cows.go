func getHint(secret string, guess string) string {
    if len(secret) != len(guess) {return ""}
    bulls := 0
    cows := 0
    freq := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
            if freq[secret[i]] < 0 {cows++}
            if freq[guess[i]] > 0 {cows++}
            freq[secret[i]]++
            freq[guess[i]]--
        }
    }
    return fmt.Sprintf("%vA%vB", bulls, cows)
}

// func getHint(secret string, guess string) string {
//     if len(secret) != len(guess) {return ""}
//     bulls := 0
//     freq := map[byte]int{}
//     for i := 0; i < len(secret); i++ {
//         if secret[i] == guess[i] {
//             bulls++
//         } else {
//             freq[secret[i]]++
//         }
//     }
//     cows := 0
//     for i := 0; i < len(secret); i++ {
//         if secret[i] == guess[i] {continue}
//         if freq[guess[i]] != 0 {
//             freq[guess[i]]--
//             cows++
//         }
//     }
//     return fmt.Sprintf("%vA%vB", bulls, cows)
// }