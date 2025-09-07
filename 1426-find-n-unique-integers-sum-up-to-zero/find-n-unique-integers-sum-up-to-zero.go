func sumZero(n int) []int {
    res := []int{}
    targetSize := n
    if n % 2 != 0 {res = append(res, 0); n--}
    left := 1
    for len(res) != targetSize {
        res = append(res, left)
        res = append(res, -left)
        left++
    }
    return res
}