func addStrings(num1 string, num2 string) string {
    n1 := len(num1)-1
    n2 := len(num2)-1
    carry := 0
    tmp := []int{}
    for n1 >= 0 || n2 >= 0 {
        n1Val := 0
        n2Val := 0
        if n1 >= 0 {n1Val = int(num1[n1]-'0'); n1--}
        if n2 >= 0 {n2Val = int(num2[n2]-'0'); n2--}
        sum := n1Val + n2Val + carry
        tmp = append(tmp, sum%10)
        carry = sum / 10        
    }
    if carry != 0 {tmp = append(tmp, carry)}
    out := new(strings.Builder)
    for i := len(tmp)-1; i >= 0; i-- {
        out.WriteString(fmt.Sprintf("%v", tmp[i]))
    }
    return out.String()
}