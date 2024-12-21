func decodeString(s string) string {
    st := []string{}
    numSt := []int{}
    curr := new(strings.Builder)
    currN := 0
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            currN = currN * 10 + int(s[i]-'0')
        } else if s[i] >= 'a' && s[i] <= 'z' {
            curr.WriteByte(s[i])
        } else if s[i] == '[' {
            st = append(st, curr.String())
            numSt = append(numSt, currN)
            currN = 0
            curr.Reset()             
        } else if s[i] == ']' {
            k := numSt[len(numSt)-1]
            numSt = numSt[:len(numSt)-1]
            decode := new(strings.Builder)
            currStr := curr.String()
            for k != 0 {
                decode.WriteString(currStr)
                k--
            }
            parent := st[len(st)-1]
            st = st[:len(st)-1]
            curr.Reset()
            curr.WriteString(parent)
            curr.WriteString(decode.String())
        }
    }
    return curr.String()
}