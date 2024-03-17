func minDays(bloomDay []int, m int, k int) int {
    left := math.MaxInt64
    right := math.MinInt64
    for i := 0; i < len(bloomDay); i++ {
        left = min(left, bloomDay[i])
        right = max(right, bloomDay[i])
    }
    ans := -1
    for left <= right{
        // how many bouquets can we make if current day was $mid?
        mid := left + (right-left)/2

        flowers := 0
        bouquets := 0
        for i := 0; i < len(bloomDay); i++ {
            // can use this flower, because it has either bloomed today or before today ( today = $mid )
            if bloomDay[i] <= mid {
                flowers++
                // have we successfully collected k adj flowers?
                if flowers == k {
                    // then we make a bouquet
                    bouquets++
                    flowers = 0
                }
            } else {
                // we cant pick this, 
                // therefore our adj counter is now broken
                // therefore reset our contagious counter ( i.e we have no adj flowers )
                flowers = 0
            }
        }

        // does $mid work?
        // $mid , i.e day works, if 
        // we have made atleast(or-more) than m bouquets
        if bouquets >= m {
            // save this as potential ans, search left
            // since we want the earliest such day
            ans = mid
            right = mid-1
        } else {
            // could not make m bouquets on $mid day
            // because we didnt have enough contagious flowers
            // therefore search right, try to use a later day
            left = mid+1
        }
    }
    return ans
}

/*
    approach: brute force
    - evaluate each day, one by one
        - from day 1
        - slight optimization: 
            - if there is no flower blooming on day 1, does it even make sense to start from day 1 ?
            - if there is no flower blooming on day 2, does it even make sense to start from day 2 ?
            - so ideally, we should start from the earliest possibe (i.e smallest bloomDay )
        - to the last possible day ( largest val in bloomDay array )
    - then for each day, we need to check whether we can make m bouquets  
        - using flowers if they have bloomed
        - that too, can only pick adj flowers (i.e contagious flowers)
        - when can we pick a flower?
            - if it has bloomed (today or before)
        - how do we keep track of how many contagiuous flowers we have picked this far?
            - counter
            - if we picked, counter++
                - and we have picked k total adj flowers this far, it means we can make a bouquet
                - therefore bouquetCounter++
            - if we couldnt, and counter is > 0, counter--
        - finally, after going thru all flowers, we can check if we successfully make m bouquets
            - if yes, exit early and return current day
            - since we are being greedy and starting from the earliest day
            - we can return and exit early as soon as a day works
            - because we want the earliest such day

    time = o(lastDay-earliestDay+1 * n)
*/

// func minDays(bloomDay []int, m int, k int) int {
//     start := math.MaxInt64
//     end := math.MinInt64
//     for i := 0; i < len(bloomDay); i++ {
//         start = min(start, bloomDay[i])
//         end = max(end, bloomDay[i])
//     }

//     for i := start; i <= end; i++ {
//         day := i
//         numBouquets := 0
//         flowersPicked := 0
//         for j := 0; j < len(bloomDay); j++ {
//             if bloomDay[j] <= day {
//                 flowersPicked++
//                 if flowersPicked == k {
//                     numBouquets++
//                     flowersPicked = 0
//                 }
//             } else {
//                 if flowersPicked > 0 {
//                     flowersPicked--
//                 }
//             }
//         }
//         if numBouquets >= m {
//             return day
//         }
//     }
//     return -1
// }


func min(x, y int) int {
    if x < y {return x}
    return y
}

func max(x, y int) int {
    if x > y {return x}
    return y
}