func candy(ratings []int) int {
    n := len(ratings)
    total := n // start by givig 1 to each child
    i := 1
    for i < n {
        for i < n && ratings[i] == ratings[i-1] {i++; continue}
        peak := 0
        for i < n && ratings[i] > ratings[i-1] {peak++; total += peak; i++}
        dip := 0
        for i < n && ratings[i] < ratings[i-1] {dip++; total += dip; i++}
        total -= min(peak, dip)
    }
    return total
}

// func candy(ratings []int) int {
//     candies := make([]int, len(ratings))
//     for i := 0; i < len(ratings); i++ {
//         candies[i] = 1
//         if i > 0 && ratings[i] > ratings[i-1] {
//             candies[i] = candies[i-1]+1
//         }
//     }
//     total := candies[len(candies)-1]
//     for i := len(ratings)-2; i >= 0; i-- {
//         if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
//             candies[i] = candies[i+1]+1
//         }
//         total += candies[i]
//     }
//     return total
// }