func judgeSquareSum(c int) bool {
    left := 0
    right := int(math.Sqrt(float64(c)))
    for left <= right {
        sum := (left*left) + (right*right)
        if sum == c {return true}
        if sum > c {
            right--
        } else {
            left++
        }
    }
    return false
}