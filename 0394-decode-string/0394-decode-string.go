func decodeString(s string) string {
    curr := new(strings.Builder)
    num := 0
    strSt := []*strings.Builder{}
    numSt := []int{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            num = num * 10 + int(char-'0')
        } else if char >= 'a' && char <= 'z' {
            curr.WriteByte(char)
        } else if char == '[' {
            strSt = append(strSt, curr)
            numSt = append(numSt, num)
            curr = new(strings.Builder)
            num = 0
        } else if char == ']' {
            times := numSt[len(numSt)-1]; numSt = numSt[:len(numSt)-1]
            parent := strSt[len(strSt)-1]; strSt = strSt[:len(strSt)-1]
            for k := 0; k < times; k++ {
                parent.WriteString(curr.String())
            }
            curr = parent
        }
    }
    return curr.String()
}