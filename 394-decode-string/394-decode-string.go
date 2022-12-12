func decodeString(s string) string {
    
    var num int
    curr := new(strings.Builder)
    strSt := []*strings.Builder{}
    numSt := []int{}
    
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            num = num * 10 + int(s[i]-'0')
        } else if s[i] == '[' {
            numSt = append(numSt, num)
            strSt = append(strSt, curr)
            num=0
            curr = new(strings.Builder)
        } else if s[i] == ']' {
            decoded := new(strings.Builder)
            k := numSt[len(numSt)-1]; numSt = numSt[:len(numSt)-1]
            for i := 0; i < k ; i++ {
                decoded.WriteString(curr.String())
            }
            curr = strSt[len(strSt)-1]
            strSt = strSt[:len(strSt)-1]
            curr.WriteString(decoded.String())
        } else {
            curr.WriteByte(s[i])
        }
    }
    return curr.String()
    
}
