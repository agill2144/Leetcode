func addStrings(num1 string, num2 string) string {
    carry := 0
    n1 := len(num1)-1
    n2 := len(num2)-1
    res := ""
    for n1 >= 0 || n2 >= 0 {
        sum := carry
        if n1 >= 0 {sum += int(num1[n1]-'0'); n1--;}
        if n2 >= 0 {sum += int(num2[n2]-'0'); n2--;}
        res = fmt.Sprintf("%v", sum % 10) + res
        carry = sum / 10
    }
    if carry != 0 {
        res = fmt.Sprintf("%v", carry) + res
        carry = 0
    }
    return res
}