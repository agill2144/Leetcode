func candy(ratings []int) int {
    n := len(ratings)
    total := n
    i := 1
    for i < n {
        for i < n && ratings[i] == ratings[i-1] {i++; continue}
        
        peak := 0
        for i < n && ratings[i] > ratings[i-1] {
            peak++
            total += peak
            i++
        }

        dip := 0
        for i < n && ratings[i] < ratings[i-1] {
            dip++
            total += dip
            i++
        }

        total -= min(peak, dip)
    }
    return total
}   