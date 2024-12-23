/*
    approach: 1 pass with freq map
    - if s[i] == guess[i] ; happy path, increment bulls counter
    - otherwise
        - we need to see if secret[i] was previously seen in guess?
        - meaning from i -> 0 in guess string, have we had a non-matching secret[i] char?
        - to build this non-matching guess char;
            - we will create a negative freq of guess[i] char; i.e freq[guess[i]]--
            - this means, we had a guess char that had not matched with a secret char
        - now to check if this secret char can match with a previously seen non-matching guess char
            - check in freq if secret[i] has a negative value?
            - negative value freq are only added by guess[i] chars when they dont match
            - so if there is a negative freq, there is a previously seen guess char thats misplaced
            - therefore cows++
            - if freq[secret[i]] < 0 {cows++}
        - similarly, we need to check if guess[i] could be paired with a non-matching previously seen secret char?
        - meaning we need to go back from i -> 0 in secret string to check if there was a non-matching guess[i] char
        - to build non-matching secret char;
            - we will create a positive freq for each secret[i] char that has not been paired with a guess char
            - now guess[i] can perform a lookup in freq to check whether secret string ever added a non matching guess[i] char
            - if there is a freq > 0, it means yes, secret string added a char that previously did not match
            - meaning we have a secret char that matches guess char but in the wrong position
            - if freq[guess[i]] > 0 {cows++}
        - now add secret freq ; freq[secret[i]]++
            - this does 2 things
            - adds itself if it did not match previously seen non-matching guess char
            - or 
            - it matched with previously seen guess char, and its freq was negative, incrementing it made it 0, therefore no-longer usable
        - now decrement guess freq; freq[guess[i]]--
            - this also does 2 things
            - it adds itself for the first time when there was no previously seen secret char to pair with
            - or
            - it matched with a previously seen secret char and now decremeting its freq because that secret char is no longer usable
        
        time = o(n)
        space = o(1) // 26 chars in alphabets
*/
func getHint(secret string, guess string) string {
    if len(secret) != len(guess) {return ""}
    bulls := 0
    cows := 0
    freq := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        s, g := secret[i], guess[i]
        if s == g {
            bulls++
        } else {
            // could s match with a previously seen g? (g is added with negative freq )
            // was there a previously seen g char that s can pair with?
            // non-matching g char have negative freq
            if freq[s] < 0 {cows++} 

            // could g match with previously seen s? (s is added with positive freq )
            // was there a previosly seen s char that g can pair with?
            // non matching s char have postive freq
            if freq[g] > 0 {cows++}
            // increment s freq 
            // either adding for the first time, or paired with previously non-matching g char and now its unusable
            freq[s]++
            // decrement g freq
            // either adding for the first time, or g paired with previously non-matchong s chart and now its unusable
            freq[g]--
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