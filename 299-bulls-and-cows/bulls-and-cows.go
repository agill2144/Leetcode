/*
    approach: 1 pass with offset freq matching
    - if both chars match, great, happy path, bulls++

    - when both chars do not match
        - they could match later tho ... 
        - secret[i] could match later
        - guess[i] could match later
        - so we need to store them in some searchable data structure
        - i.e hash map
    - this hash map is tracking unmatched chars ONLY
    - when secret[i] exists in hash map 
        - how do we know whether it was added by guess string?
    - when guess[i] exists in hash map
        - how do we know whether it was added by secret string?
    - so then we should create a partition from where chars are added from
        - {secretChars: {}, guessChars: {}}
        - or 2 separate arrays tracking freq of each char for each string
            upto an idx i ( curr i position )
    - now when we have an unmatched char, then
    - we check if guess char exists in secret partition
        - did secret string before current idx have this guess[i]char?
        - if yes, cow detected, decrement its count in secret partition space
        - because its no longer usable and its matched 
        - if not, we need to store this guess char to see if it can be matched later
    - we need to check other way around too
    - if secret char exists in guess partition
        - did the guess string before current idx have this secret[i] char?
        - if yes, cow detected, decrement its count in guess partition space
        - because its no longer usable and its matched 
        - if not, we need to store this secret char to see if it can be matched later


        time = o(n)
        space = o(2*10) = o(1) 
*/
func getHint(secret string, guess string) string {
    if len(secret) != len(guess) { return "" }

    bulls := 0
    cows := 0
    // freq := map[string]map[byte]int{"s": {}, "g": {}}
    sFreq := make([]int, 10)
    gFreq := make([]int, 10)

    for i := 0; i < len(secret); i++ {
        s, g := secret[i], guess[i]
        if s == g {
            bulls++
        } else {
            // have we seen g char in secret before this position?
            if sFreq[int(g-'0')] > 0 {
                // if yes, cow detected
                cows++
                sFreq[int(g-'0')]--
            } else {
                // store this guess char to match later
                gFreq[int(g-'0')]++
            }

            // have we seen s char in guess string before this idx?
            if gFreq[int(s-'0')] > 0 {
                // if yes, cow detected
                cows++
                gFreq[int(s-'0')]--
            } else {
                // store this secret char to match later
                sFreq[int(s-'0')]++
            }
        }
    }

    return fmt.Sprintf("%dA%dB", bulls, cows)
}



// 2 pass, freq map
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