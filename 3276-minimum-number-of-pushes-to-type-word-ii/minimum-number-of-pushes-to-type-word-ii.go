type pair struct {
    char byte
    freq int
}

func minimumPushes(word string) int {
    freq := map[byte]int{}
    for i := 0; i < len(word); i++ {freq[word[i]]++}
    freqArr := []pair{}
    for k, v := range freq {freqArr = append(freqArr, pair{k, v})}
    sort.Slice(freqArr, func(i, j int)bool{
        return freqArr[i].freq > freqArr[j].freq
    })
    
    keyPad := map[int][]byte{}
    digit := 2
    total := 0
    for i := 0; i < len(freqArr); i++ {
        if digit > 9 {digit = 2}
        keyPad[digit] = append(keyPad[digit], freqArr[i].char)
        times := len(keyPad[digit])
        total += (freqArr[i].freq * times)
        digit++
    }

    return total
}