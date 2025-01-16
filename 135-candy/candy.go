func candy(ratings []int) int {
    n := len(ratings)
    total := n
    i := 1
    for i < len(ratings) {
        for  i < len(ratings) && ratings[i] ==ratings[i-1] {i++; continue}
        if i == len(ratings) {break}
        peak := 0
        valley := 0
        for i < len(ratings) && ratings[i] > ratings[i-1] {peak++; i++}
        for i < len(ratings) && ratings[i] < ratings[i-1] {valley++; i++}
        peakTotal := (peak*(peak+1))/2 
        valleyTotal := (valley*(valley+1))/2
        total += (peakTotal + valleyTotal - min(peak, valley)) 
    }
    return total
}