func decodeString(s string) string {
    numSt := []int{}
    parentSt := []*strings.Builder{}
    
    num := 0
    curr := new(strings.Builder)
    
    for i := 0 ; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            num = num * 10 + int(char-'0')
        } else if char == '[' {
            numSt = append(numSt, num)
            parentSt = append(parentSt, curr)
            curr = new(strings.Builder)
            num = 0
        } else if char == ']' {
            nTimes := numSt[len(numSt)-1]; numSt = numSt[:len(numSt)-1]
            tmpBuilder := new(strings.Builder)
            for i := 0; i < nTimes; i++ {
                tmpBuilder.WriteString(curr.String())
            }
            parent := parentSt[len(parentSt)-1]; parentSt = parentSt[:len(parentSt)-1]
            parent.WriteString(tmpBuilder.String())
            curr = parent
        } else {
            curr.WriteByte(char)
        }
    }
    
    return curr.String()
}