func candy(ratings []int) int {
    n := len(ratings)
    total := n
    peakSum := 0
    valleySum := 0
    i := 1
    for i < len(ratings) {
        for i < len(ratings) && ratings[i] == ratings[i-1] {
            i++
            continue
        }
        if i == len(ratings) {break}

        for i < len(ratings) && ratings[i] > ratings[i-1] {
            peakSum++
            i++
        }
        for i < len(ratings) && ratings[i] < ratings[i-1] {
            valleySum++
            i++
        }

        t := ((peakSum*(peakSum+1))/2) + ((valleySum*(valleySum+1))/2)
        t -= min(peakSum,valleySum)
        total += t
        peakSum = 0
        valleySum = 0
    }
    return total
}