
func getNoZeroIntegers(n int) []int {
    for i := 1; i < n; i++ {
        right := n-i
        leftStr := fmt.Sprintf("%v", i)
        rightStr := fmt.Sprintf("%v", right)
        if !strings.Contains(leftStr, "0") && !strings.Contains(rightStr, "0") {
            return []int{i, right}
        }
    }
    return nil
}
