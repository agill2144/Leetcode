/*
    approach: brute force ( does not work )
    - conver to int
    - sum
    - return str(sum)
    - integer overflow!
    - can we use float64?
        - yes and no
        - yes, because conversions might work
        - but the result might be 1.123e^19
        - which does not match the format of output

    approach: paper math
    - add 2 digits at a time, no worry about overflow
    - but you do have to worry about carry
        - 9+9 = 18
        - carry 1 over
    - paper math, we start from the right to left
    - we will do the same
    - we will first write the sum of each digit from right to left
        - in a temp array
    - then we will loop from back of the temp array to create the
        final output string
    - we started summing from the back, therefore the sum is saved backwards
    - hence the need to loop from back of the array


*/
func addStrings(num1 string, num2 string) string {
    tmp := []int{}
    n1 := len(num1)-1
    n2 := len(num2)-1
    carry := 0
    for n1 >= 0 || n2 >= 0 {
        n1Val := 0
        if n1 >= 0 {n1Val = int(num1[n1]-'0'); n1--}
        n2Val := 0
        if n2 >= 0 {n2Val = int(num2[n2]-'0'); n2--}
        sum := n1Val + n2Val + carry
        tmp = append(tmp, sum%10)
        carry = sum / 10
    }
    if carry != 0 {
        tmp = append(tmp, carry)
        carry = 0
    }
    res := new(strings.Builder)
    for i := len(tmp)-1; i >= 0; i-- {
        res.WriteString(fmt.Sprintf("%v", tmp[i]))
    }
    return res.String()
}