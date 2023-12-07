// largest odd number will be the first odd digit we run into from the right side
func largestOddNumber(num string) string {
    for i := len(num)-1; i >= 0; i-- {
        if int(num[i]-'0') % 2 != 0 {return num[:i+1]}
    }
    return ""
}