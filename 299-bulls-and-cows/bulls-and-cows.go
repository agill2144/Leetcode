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


    EDGE CASES
    Why cant we create a freq map and check for bulls and cows later in 2nd pass together?
    - secret = 1122 -> freq = {1:2, 2:2}
    - guess  = 1222
    - we end up counting extra cows because
    - we have 3 occurence of 2s in guess
    - but only have 2 occurence in freq map
    - when i is at idx 1, we see char mismatch
    - then we will check if freq map contains our guess[i] char
    - and yes, it does, and therefore cows++ and freq[guess[i]]--
    - when i is done at idx 1 , freq = {1:2, 2:1}
    - when i is done at idx 2 , freq = {1:2, 2:0}
    - when i is get at idx 3,
        - this is not a char mismatch
        - therefore bulls++
        - and this char freq needs to be decremented, as its curr instance can no longer be used
        - but doing so will make freq negative
        - when freq goes negative for a secret char, it means
        - that we used a occurence that was extra in guess to match with this char in secret
        - meaning somewhere we did cows++, and therefore thats extra 
        - to fix this, we can check if freq goes negative and if yes, 
        - it means a extra guess char exists that was used to match with secret char
        - as a cow char, therefore decrement cows by 1
    - alternatively we could do this in 2 passes,
    - count bulls first when creating freq map
    - then count cows
    - if there were perfect char matches, their occurence never got placed into freq map
    - therefore occurences that exist in freq map are purely for unmatched chars
    - and can only be used cows
    

    
    n = len(guess || secret)
    tc = o(2n) = o(n)
    sc = o(1) = freq map stores digits from 0 to 9, therefore constant

*/
func getHint(secret string, guess string) string {
    bulls := 0
    cows := 0
    freq := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        freq[secret[i]]++
    }
    for i := 0; i < len(guess); i++ {
        if guess[i] == secret[i] {
            bulls++
            freq[guess[i]]--
            if freq[guess[i]] < 0 {cows--}
        } else if freq[guess[i]] > 0 {
            cows++
            freq[guess[i]]--
        }
    }
    return fmt.Sprintf("%vA%vB",bulls, cows)
}
