func decodeString(s string) string {
    parentSt := []*strings.Builder{}
    numSt := []int{}
    num := 0
    curr := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            curr.WriteByte(s[i])
        } else if s[i] >= '0' && s[i] <= '9' {
            num = num * 10 + int(s[i]-'0')
        } else if s[i] == '[' {
            parentSt = append(parentSt, curr)
            numSt = append(numSt, num)
            curr = new(strings.Builder)
            num = 0
        } else if s[i] == ']' {
            k := numSt[len(numSt)-1]
            numSt = numSt[:len(numSt)-1]
            tmp := new(strings.Builder)
            for k != 0 {
                tmp.WriteString(curr.String())
                k--
            }
            parentStr := parentSt[len(parentSt)-1]
            parentSt = parentSt[:len(parentSt)-1]
            parentStr.WriteString(tmp.String())
            curr = parentStr
        }
    }
    return curr.String()
}
