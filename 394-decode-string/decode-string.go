func decodeString(s string) string {
    strSt := []*strings.Builder{}
    numSt := []int{}
    str := new(strings.Builder)
    num := 0
    for i := 0; i < len(s); i++ {
        curr := s[i]
        if curr >= 'a' && curr <= 'z' {
            str.WriteByte(curr)
        } else if curr >= '0' && curr <= '9' {
            num = num * 10 + int(curr-'0')
        } else if curr == '[' {
            strSt = append(strSt, str)
            numSt = append(numSt, num)
            str = new(strings.Builder)
            num = 0
        } else if curr == ']' {
            // repeat curr string topNum of times
            topN := numSt[len(numSt)-1]
            numSt = numSt[:len(numSt)-1]
            tmp := new(strings.Builder)
            for i := 0; i < topN; i++ {
                tmp.WriteString(str.String())
            }
            // combine with parent
            parent := strSt[len(strSt)-1]
            strSt = strSt[:len(strSt)-1]
            parent.WriteString(tmp.String())
            // make curr str as parent 
            str = parent
        }
    }
    return str.String()
}