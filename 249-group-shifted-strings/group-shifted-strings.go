/*
    observartion:
    - 2 words are groupped together when
        - word1 is of same size as word2
        - word1 has the same shifting seq as word2
            - diff between each contagious char in word1 == diff between each contagious char in word2
    
    approach: hashing
    - we will hash each word to a hash key
    - hash key will be the difference between each contagious char
        - example: "abc" -> 1-1
        - the diff between b-a is 1
        - the diff between c-b is 1
        - therefore the hash key is 1-1
        - add delimeter to avoid key clashes
    - the diff could go negative!!!
        - when diff goes negative, add a +26    
    - what about single length words?
        - well all single len words belong to 1 group
        - a, b, c, ...z all belong to 1 group
        - so if word size is 1
        - then our hash function can return a static const key
        - lets call this const key "0"
    - loop over all the words one by one
    - create a hash of each word
    - group the words with same hash key in a hashmap
    - then take another pass to collect similar words that belong to the same group / key

    tc = o(n*k) + o(n)
    sc = o(n)

    another hashing function that can be done is;
    - instead of creating a hash key with differences between contagious char
    - take the difference 
    - and get the alphabet that exist at that diff idx
    - the diff could go negative!!!
        - example: az, ba
        - processing "ba" and at a char
        - a-b = will produce a negative diff
        - -1, to make it correct and relative to
        - idxs available in alphabets, add a +26
        - -1+26 = 25
        - now get the char at idx 25 = z
        - processing "az" and at z char
        - z-a = 25
        - now get the char at idx 25 = z
        - therefore az, ba produce the same hashkey

*/
func groupStrings(strings []string) [][]string {
    grps := map[string][]string{}
    for i := 0; i < len(strings); i++ {
        key := hash(strings[i])
        grps[key] = append(grps[key], strings[i])
    }
    out := [][]string{}
    for _, v := range grps {
        out = append(out,v)
    }
    return out
}

func hash(word string) string {
    if len(word) <= 1 {return "0"}
    out := new(strings.Builder)
    for i := 1; i < len(word); i++ {
        curr := int(word[i]-'a')
        prev := int(word[i-1]-'a')
        diff := curr-prev
        if diff < 0 {diff += 26}
        out.WriteString(fmt.Sprintf("%v-", diff))
    }
    return out.String()
}