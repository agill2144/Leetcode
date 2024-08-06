type pair struct {
    char byte
    freq int
}

/*
    greedy, assign the highest freq char 
    to be the first char on an empty key in the keypad
    
    word = "accccccccccccbbb"
    because we are allowed to remap any char to any key
    we should make sure that c and b are the first char on a key
    so that we dont have to press a key more than once to get all occurences of c and b
    therefore highest repeating char is assigned to an empty key if possible

    w = len(word)
    
    time;
    o(w) + o(26 log 26) + o(26) = o(w)
    
    space;
    o(26) + o(26) + o(9 digits in keyPad) = o(1)
*/
func minimumPushes(word string) int {
    freq := map[byte]int{}
    for i := 0; i < len(word); i++ {freq[word[i]]++} 
    freqArr := []pair{}
    for k, v := range freq {freqArr = append(freqArr, pair{k, v})}
    sort.Slice(freqArr, func(i, j int)bool{
        return freqArr[i].freq > freqArr[j].freq
    })
    
    keyPad := map[int]int{} 
    digit := 2
    total := 0
    for i := 0; i < len(freqArr); i++ {
        if digit > 9 {digit = 2}
        keyPad[digit]++
        times := keyPad[digit]
        total += (freqArr[i].freq * times)
        digit++
    }

    return total
}