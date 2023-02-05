func decodeString(s string) string {
    charSt := []*strings.Builder{}
    numSt := []int{}
    
    char := new(strings.Builder)
    num := 0
    
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c >= '0' && c <= '9' {
            num = num * 10 + int(s[i]-'0')
        } else if c >= 'a' && c <= 'z' {
            char.WriteByte(c)
        } else if c == '[' {
            charSt = append(charSt, char)
            char = new(strings.Builder)
            numSt = append(numSt, num)
            num = 0
        } else if c == ']' {
            times := numSt[len(numSt)-1]
            numSt = numSt[:len(numSt)-1]
            decodedStr := new(strings.Builder)
            for i := 0; i < times; i++ {
                decodedStr.WriteString(char.String())
            }
            parentChar := charSt[len(charSt)-1]
            charSt = charSt[:len(charSt)-1]
            parentChar.WriteString(decodedStr.String())
            char = parentChar
        }
    }
    return char.String()
}