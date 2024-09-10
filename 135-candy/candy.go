func candy(ratings []int) int {
    n := len(ratings)
    candies := make([]int,n)
    candies[0] = 1
    for i := 1; i < n; i++ {
        candies[i] = 1
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1]+1
        }
    }
    total := candies[n-1]
    for i := n-2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1]{
            candies[i] = candies[i+1]+1
        }
        total += candies[i]
    }
    return total
}

// func candy(ratings []int) int {
//     n := len(ratings)
//     total := n
//     i := 1
//     for i < n {
//         for i < n && ratings[i] == ratings[i-1] {i++; continue}
        
//         peak := 0
//         for i < n && ratings[i] > ratings[i-1] {
//             peak++
//             total += peak
//             i++
//         }

//         dip := 0
//         for i < n && ratings[i] < ratings[i-1] {
//             dip++
//             total += dip
//             i++
//         }

//         total -= min(peak, dip)
//     }
//     return total
// }   