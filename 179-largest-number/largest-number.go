func largestNumber(nums []int) string {
    numStr := []string{}
    for i := 0; i < len(nums); i++ {
        numStr = append(numStr, fmt.Sprintf("%v",nums[i]))
    }
    sort.Slice(numStr, func(i, j int)bool{
        iVal := numStr[i]
        jVal := numStr[j]
        ij := iVal + jVal
        ji := jVal + iVal
        if strings.Compare(ij, ji) == 1 {
            return true
        } else {
            return false
        }
    })
    out := new(strings.Builder)
    for i := 0; i < len(numStr); i++ {
        if len(out.String()) == 0 && numStr[i] == "0" {continue} 
        out.WriteString(numStr[i])
    }
    if out.String() == "" {return "0"}
    return out.String()
}