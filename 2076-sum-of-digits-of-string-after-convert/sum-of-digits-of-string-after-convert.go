func getLucky(s string, k int) int {
    str := ""
    for i := 0; i < len(s); i++ {
        str += fmt.Sprintf("%v",int(s[i]-'a')+1)
    }

    for k != 0 {
        rSum := 0
        for i := 0; i < len(str); i++ {
            dig := int(str[i]-'0')
            rSum += dig
        }
        k--
        if k == 0 {return rSum}
        str = fmt.Sprintf("%v", rSum)
    }
    return -1
}