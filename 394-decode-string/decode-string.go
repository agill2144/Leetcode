func decodeString(s string) string {
    strSt := []string{}
    numSt := []int{}
    currNum := 0
    currStr := ""
    for i := 0; i < len(s); i++ {
        if s[i] == '[' {
            numSt = append(numSt, currNum)
            strSt = append(strSt, currStr)
            currNum = 0
            currStr = ""
        } else if s[i] == ']' {
            topNum := numSt[len(numSt)-1]
            numSt = numSt[:len(numSt)-1]
            tmp := new(strings.Builder)
            for k := 0; k < topNum; k++ {
                tmp.WriteString(currStr)
            }
            parentStr := strSt[len(strSt)-1]
            strSt = strSt[:len(strSt)-1]
            currStr = parentStr + tmp.String()
        } else if s[i] >= '0' && s[i] <= '9' {
            currNum = currNum * 10 + int(s[i]-'0')
        } else {
            currStr += string(s[i])
        }
    }
    return currStr
}