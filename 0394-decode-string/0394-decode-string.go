func decodeString(s string) string {
    numSt := []int{}
    strSt := []*strings.Builder{}
    curr := new(strings.Builder)
    num := 0
    
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            num = num * 10 + int(s[i]-'0')
        } else if s[i] == '[' {
            strSt = append(strSt, curr)
            numSt = append(numSt, num)
            num = 0
            curr = new(strings.Builder)
        } else if s[i] == ']' {
            topNum := numSt[len(numSt)-1]; numSt = numSt[:len(numSt)-1]
            flattend := new(strings.Builder)
            for i := 0; i < topNum; i++ {
                flattend.WriteString(curr.String())
            }
            curr = strSt[len(strSt)-1]; strSt = strSt[:len(strSt)-1]
            curr.WriteString(flattend.String())
        } else {
            curr.WriteByte(s[i])
        }
    }
    return curr.String()
}