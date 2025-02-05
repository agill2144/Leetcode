/*
    approach: freq map and two-ptrs
    - inorder to match guess char with secret char
    - we need to know whether guess char exists in secret char
    - its also possible that a guess char repeats more than 1 time
    - therefore we need to check the same guess char exists in secret that many times
    - hence we will store each char in secret string into a freq map to make searching easy
    - while building the freq map , we can also start counting bulls
    - i.e chars that are correctly positioned 
    - if guess[i] == secret[i], these chars are correctly positioned
    - then we dont need to include secret[i] occurence in freq match
        because it has fully matched with correct position in guess secret
    - finally, we can take another pass, counting all the cows
    - i.e if secret[i] != guess[i] and if the guess char exists in secret freq
    - then we can increment cows counter and reduce this char freq from secret freq
        because its current occurence can no longer used to match later
    
    n = len(guess || secret)
    tc = o(2n) = o(n)
    sc = o(1) = freq map stores digits from 0 to 9, therefore constant

*/
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
