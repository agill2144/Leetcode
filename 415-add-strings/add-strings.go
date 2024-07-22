func addStrings(num1 string, num2 string) string {
    tmp := []int{}
    n1 := len(num1)
    n2 := len(num2)
    carry := 0
    p1 := n1-1
    p2 := n2-1
    for p1 >= 0 || p2 >= 0 {
        p1Val := 0
        if p1 >= 0 {p1Val = int(num1[p1]-'0'); p1--}
        p2Val := 0
        if p2 >= 0 {p2Val = int(num2[p2]-'0'); p2--}

        total := p1Val + p2Val + carry
        toAdd := total%10
        carry = total/10
        tmp = append(tmp, toAdd)
    }
    if carry != 0 {
        tmp = append(tmp, carry)
    }
    out := new(strings.Builder)
    for i := len(tmp)-1; i >= 0; i-- {
        out.WriteString(fmt.Sprintf("%v", tmp[i]))
    }
    return out.String()
}

// 24 / 10 = 2
// 2753 / 10 = 275