func candy(ratings []int) int {
    total := len(ratings)
    i := 1
    for i < len(ratings) {
        for i < len(ratings) && ratings[i] == ratings[i-1] {i++}
        peak := 0
        for i < len(ratings) && ratings[i] > ratings[i-1] {
            peak++
            total += peak
            i++
        }
        dip := 0
        for i < len(ratings) && ratings[i] < ratings[i-1] {
            dip++
            total += dip
            i++
        }
        total -= min(peak, dip)
    }
    return total
}