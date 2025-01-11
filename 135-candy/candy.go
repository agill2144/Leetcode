func candy(ratings []int) int {
    n := len(ratings)
    count := make([]int, n)
    for i := 0; i < len(ratings); i++ {
        count[i] = 1
        curr := ratings[i]
        prev := math.MaxInt64
        if i-1 >= 0 {prev = ratings[i-1]}
        if curr > prev {
            count[i] = count[i-1]+1
        }
    }
    total := 0
    for i := n-1; i >= 0; i-- {
        curr := ratings[i]
        prev := math.MaxInt64
        if i+1 < n {prev = ratings[i+1]}
        if curr > prev && count[i] <= count[i+1] {count[i] = count[i+1]+1}
        total += count[i]
    }
    return total
}