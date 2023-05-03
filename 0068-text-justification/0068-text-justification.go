func fullJustify(words []string, maxWidth int) []string {
    out := []string{}
    i := 0
    n := len(words)
    for i < n {
        lineLen := len(words[i])
        j := i+1
        for j < n {
            numWords := j-i+1
            minSpaces := numWords-1
            newLineLen := lineLen + len(words[j]) + minSpaces
            if newLineLen > maxWidth {break}
            lineLen += len(words[j])
            j++
        }
        diff := maxWidth-lineLen
        numWords := j-i
        // if its the last line OR number of words in this line is only 1, then we add all spaces to the right
        if numWords == 1 || j >= n {
            out = append(out, addSpacesToRight(i, j, words, diff))
        // otherwise spread out the spaces evenly as possible, and whatever is left after even distribution, left will be assigned more spaces than the slots on the right
        } else {
            out = append(out, addSpacesToMiddleEvenly(i,j, words, diff))
        }
        i = j        
    }
    return out
}

func addSpacesToRight(i, j int, words []string, diff int) string {
    tmp := new(strings.Builder)
    remainingSpaces := diff - (j-i-1)
    for k := i; k < j; k++ {
        tmp.WriteString(words[k])
        if k < j-1 {
            tmp.WriteString(" ")
        }
    }
    for remainingSpaces != 0 {tmp.WriteString(" "); remainingSpaces--}
    return tmp.String()
}

func addSpacesToMiddleEvenly(i, j int, words []string, diff int) string {
    tmp := new(strings.Builder)
    numSpaces :=  j-i-1
    spaceSize := diff/numSpaces
    extraSpaces := diff%numSpaces
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