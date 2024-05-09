func maximumHappinessSum(happiness []int, k int) int64 {

    n := len(happiness)
    sort.Ints(happiness)
    if k == 0 {return 0}
    if k == 1 {return int64(happiness[n-1])}

    var total int64
    i := n-1
    decreaseBy := 0
    for k > 0 && i >= 0 {
        val := happiness[i]
        val -= decreaseBy
        if val < 0 {break}
        decreaseBy++
        total += int64(val)
        k--
        i--
    }
    return total
}