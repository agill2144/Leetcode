func decodeString(s string) string {
    str := []*strings.Builder{}
    numSt := []int{}
    
    curr := 0
    currStr := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            curr = curr * 10 + int(s[i]-'0')
        } else if s[i] >= 'a' && s[i] <= 'z' {
            currStr.WriteByte(s[i])
        } else if s[i] == '[' {
            str = append(str, currStr)
            numSt = append(numSt, curr)
            curr = 0
            currStr = new(strings.Builder)
        } else if s[i] == ']' {
            k := numSt[len(numSt)-1]
            numSt = numSt[:len(numSt)-1]
            tmp := new(strings.Builder)
            for i := 0; i < k; i++ {
                tmp.WriteString(currStr.String())
            }
            parent := str[len(str)-1]
            str = str[:len(str)-1]
            parent.WriteString(tmp.String())
            currStr = parent
        }
    }
    return currStr.String()
}