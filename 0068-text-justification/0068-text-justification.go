func fullJustify(words []string, maxWidth int) []string {
    out := []string{}
    i := 0
    n := len(words)
    for i < n {
        charCount := len(words[i])
        j := i+1
        for j < n {
            numWords := j-i+1 // including this word
            minSpaces := numWords-1 // min number of spaces (1) will always be within these words, therefore numWords-1
            newCharCount := charCount + len(words[j]) + minSpaces // get the total char count with min spaces
            if newCharCount > maxWidth {break} // break out if we exceed the maxWidth
            charCount += len(words[j]) // otherwise, keep on adding the number of chars ( without min space )
            j++
        }
        // calc the remaining char count left if we have maxWidth and charCount
        remainingCount := maxWidth-charCount
        // number of words does not include the jth word, because we probably broke at that word
        numWords := j-i
        // there are 2 main cases here
        // 1. if number of words == 1 or we are on the last line; then use min space in middle and remaining spaces will be added to the end.
        // 2. otherwise we need to evenly spread our spaces
        
        // if its the last line OR number of words in this line is only 1, then we add all spaces to the right
        if numWords == 1 || j >= n {
            out = append(out, addSpacesToRight(i, j, words, remainingCount))
        // otherwise spread out the spaces evenly as possible, and whatever is left after even distribution, left will be assigned more spaces than the slots on the right
        } else {
            out = append(out, addSpacesToMiddleEvenly(i,j, words, remainingCount))
        }
        // move our i ptr to j
        i = j        
    }
    return out
}

func addSpacesToRight(i, j int, words []string, remainingCount int) string {
    tmp := new(strings.Builder)
    remainingSpaces := remainingCount - (j-i-1)
    for k := i; k < j; k++ {
        tmp.WriteString(words[k])
        if k < j-1 {
            tmp.WriteString(" ")
        }
    }
    for remainingSpaces != 0 {tmp.WriteString(" "); remainingSpaces--}
    return tmp.String()
}

func addSpacesToMiddleEvenly(i, j int, words []string,remainingCount int) string {
    tmp := new(strings.Builder)
    numSpaces :=  j-i-1
    spaceSize := remainingCount/numSpaces
    extraSpaces := remainingCount%numSpaces
    for k := i; k < j; k++ {
        tmp.WriteString(words[k])
        totalSpaces := spaceSize
        if extraSpaces > 0 {
            totalSpaces++
            extraSpaces--
        }
        if k < j-1 {
            for l := 0; l < totalSpaces; l++ {tmp.WriteString(" ")}
        }
    }
    return tmp.String()
}