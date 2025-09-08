
func getNoZeroIntegers(n int) []int {
    for i := 1; i < n; i++ {
        left := i
        right := n-i
        leftStr := fmt.Sprintf("%v", left)
        rightStr := fmt.Sprintf("%v", right)
        if !strings.Contains(leftStr, "0") && !strings.Contains(rightStr, "0") {
            return []int{left, right}
        }
    }
    return nil
}
