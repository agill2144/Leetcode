func decodeString(s string) string {
    curr := new(strings.Builder)
    n := 0
    currSt := []*strings.Builder{}
    nSt := []int{}
    
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= 'a' && char <= 'z' {
            curr.WriteByte(char)
        } else if char >= '0' && char <= '9' {
            n = n * 10 + int(char-'0')
        } else if char == '[' {
            currSt = append(currSt, curr)
            nSt = append(nSt, n)
            curr = new(strings.Builder)
            n = 0
        } else if char == ']' {
            expanded := new(strings.Builder)
            kTimes := nSt[len(nSt)-1]; nSt = nSt[:len(nSt)-1]
            for j := 0; j < kTimes; j++ {
                expanded.WriteString(curr.String())
            }
            parent := currSt[len(currSt)-1]; currSt = currSt[:len(currSt)-1]
            parent.WriteString(expanded.String())
            curr = parent
        }
    }
    return curr.String()
}