func decodeString(s string) string {
    parent :=  []*strings.Builder{}
    num := []int{}
    currNum := 0
    currStr := new(strings.Builder)
    
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
           currNum = currNum * 10 + int(s[i]-'0')
        } else if s[i] >= 'a' && s[i] <= 'z' {
            currStr.WriteByte(s[i])
        }else if s[i] == '[' {
            num = append(num, currNum)
            parent = append(parent, currStr)
            currNum = 0
            currStr = new(strings.Builder)
        } else if s[i] == ']' {
            k := num[len(num)-1]
            num = num[:len(num)-1]
            decoded := new(strings.Builder)
            for k != 0 {
                decoded.WriteString(currStr.String())
                k--
            }
            p := parent[len(parent)-1]
            parent = parent[:len(parent)-1]
            p.WriteString(decoded.String())
            currStr = p
        }
    }
    return currStr.String()
}