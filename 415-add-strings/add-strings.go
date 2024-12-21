func addStrings(num1 string, num2 string) string {
    carry := 0
    res := ""
    n1, n2 := len(num1)-1, len(num2)-1
    for n1 >= 0 || n2 >= 0 {
        n1Val := 0
        n2Val := 0
        if n1 >= 0 {n1Val = int(num1[n1]-'0'); n1--}
        if n2 >= 0 {n2Val = int(num2[n2]-'0'); n2--}
        sum := n1Val + n2Val + carry
        res = fmt.Sprintf("%v", sum%10) + res
        carry = sum / 10
    }
    if carry != 0 {
        res = fmt.Sprintf("%v", carry) + res
        carry = 0
    }
    return res
}