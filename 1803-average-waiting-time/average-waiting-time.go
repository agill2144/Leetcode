func averageWaitingTime(customers [][]int) float64 {
    var total float64 = 0.0
    var curr float64 = 0.0
    for i := 0; i < len(customers); i++ {
        a := float64(customers[i][0])
        p := float64(customers[i][1])
        overlap := curr-a
        if overlap < 0 {overlap = 0}
        total += overlap
        curr = max(curr, a)
        total += p
        curr += p
    }
    return total / float64(len(customers))

}