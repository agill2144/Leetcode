func sumZero(n int) []int {
    // if n == 1 {return []int{0}}
    // if n == 2 {return []int{1,-1}}
    num := 1
    left := (n-1)/2
    if n % 2 != 0 { left-- }
    res := make([]int, n)
    for left >= 0 {
        res[left] = num * -1
        res[n-1-left] = num
        num++
        left--
    }
    return res
}