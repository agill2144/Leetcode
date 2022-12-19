func decodeString(s string) string {
    numSt := []int{}
    strSt := []*strings.Builder{}
    num := 0
    curr := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            num = num * 10 + int(s[i]-'0')
        } else if s[i] == '[' {
            numSt = append(numSt, num)
            strSt = append(strSt, curr)
            curr = new(strings.Builder)
            num = 0
        } else if s[i] == ']' {
            k := numSt[len(numSt)-1]; numSt = numSt[:len(numSt)-1]
            decoded := new(strings.Builder)
            for j := 0; j < k; j++ {
                decoded.WriteString(curr.String())
            }
            curr = strSt[len(strSt)-1]; strSt = strSt[:len(strSt)-1]
            curr.WriteString(decoded.String())
        } else {
            curr.WriteByte(s[i])
        }
    }
    return curr.String()
}