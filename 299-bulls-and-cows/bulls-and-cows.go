// do this in few passes
/*
    1. maintain a freq map of secret string
    2. find out all the bulls
        - i.e count guess chars that are in the correct position as secret chars
        - loop over guess string, and check if guess[i] == secret[i]
        - if it does, increment bulls counter and reduce the freq of this char from freqMap since we can no longer use it
    3. find out all the cows
        - i.e count all the chars from guess that exist in secret
        - HOWEVER we should avoid double counting if char at ith idx is same (guess[i] == secret[i])
            - because our step #2 did this already, so they should not be checked again
        - if guess[i] char exists in freqMap and its count > 0
            - cows++
            - reduce freq of char in freqMap
    time = o(n) + o(n) + o(n) = o(n)
    space = 0-9 are the keys; o(n)
*/
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
        if guess[i] == secret[i] {continue}
        count, ok := digitCount[guess[i]]
        if !ok {continue}
        if count > 0 {
            cows++
            digitCount[guess[i]]--
        }
    }
    return fmt.Sprintf("%vA%vB", bulls, cows)
}