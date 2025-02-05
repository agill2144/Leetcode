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
    - now when we have an unmatched char, then
    - we check if guess char exists in secret partition
        - did secret string before current idx have this guess[i]char?
        - if yes, cow detected, decrement its count in secret partition space
        - if not, we need to store this guess char to see if it can be matched later
    - we need to check other way around too
    - if secret char exists in guess partition
        - did the guess string before current idx have this secret[i] char?
        - if yes, cow detected, decrement its count in guess partition space
        - if not, we need to store this secret char to see if it can be matched later


        time = o(n)
        space = o(1) 
        - digits are from 0-9 , so 10 possible keys at worst
        - 2 * 10 = 20 keys in each partition space
*/
func getHint(secret string, guess string) string {
    if len(secret) != len(guess) { return "" }

    bulls := 0
    cows := 0
    freq := map[string]map[byte]int{"s": {}, "g": {}}

    for i := 0; i < len(secret); i++ {
        s, g := secret[i], guess[i]
        
        if s == g {
            bulls++
        } else {
            // Check if 'g' exists in secret's storage → cow found!
            if freq["s"][g] > 0 {
                cows++
                freq["s"][g]-- // Decrement since we matched it
            } else {
                freq["g"][g]++ // Store unmatched guess character
            }

            // Check if 's' exists in guess's storage → cow found!
            if freq["g"][s] > 0 {
                cows++
                freq["g"][s]-- // Decrement since we matched it
            } else {
                freq["s"][s]++ // Store unmatched secret character
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