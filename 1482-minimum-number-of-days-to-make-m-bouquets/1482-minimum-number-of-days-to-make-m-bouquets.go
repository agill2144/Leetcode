/*
    approach: brute force
    - we want the earliest possible day, when we can make m bouquets with k adjacent flowers
    - therefore we can be greedy and start from day 1
        - better = start from the smallest possible day
    - and go up to the last possible day when flowers are bloomed
    - and then check how many adjacent flowers make bouquets
    
    time;
    o(n) ; to find min and max day
    o((max-min) * n) looping from min->max and within each iteration we are going over the array to count and check
    therefore ; o(n) + (maxVal * n)
    
    space;
    o(1)
    
    approach: binary search on answers
    - in the above brute force, we see that we have a range!
    - we are looking for an answer in a range of days ( earliest -> last possible day)
    - this is a sorted range
    - therefore the outter loop which took o(max-min) time, could be trimmed down to o(log(max-min))
    - the inner for loop remains the same
    
    time;
    o(n) ; to find left and right ptrs ( start and end of our potential ans range )
    + o(log(max-min) * n)
    
    space; o(1)
*/
func minDays(bloomDay []int, m int, k int) int {
    left := math.MaxInt64
    right := math.MinInt64
    
    // time = o(n)
    for i := 0; i < len(bloomDay); i++ {
        if bloomDay[i] < left {
            left = bloomDay[i]
        }
        if bloomDay[i] > right {
            right = bloomDay[i]
        }
    }
    
    ans := -1
    // time = o(log$maxBloomDay * n)
    for left <= right {
        mid := left + (right-left) / 2
        
        // count adjacent flowers that have bloomed on/before this day(mid)
        // and keep track of num of bouquets we were able to make on this day(mid)
        count := 0
        numBouquets := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                // this flow is bloomed
                count++
                if count == k {
                    count = 0
                    numBouquets++
                }
            } else {
                count = 0
            }
        }
        
        if numBouquets < m {
            // if we could not even make m bouquets,
            // it means we were too early on guessing which day was the correct day
            // as in, the day we picked, didnt have enough adjacent flowers blooming
            // therefore we need to pick a later day
            // therefore left = mid+1
            left = mid+1
        } else {
            // if we were able to make atleast m bouquets,
            // going to a later day after $midth day does not make any sense, because for sure they will be bloomed
            // because we are binary searching on a sorted range of days ( minVal -> maxVal )
            // save this as a potential ans, and see if we can find an earlier day instead
            // therefore right = mid-1
            ans = mid
            right = mid-1
        }
    }
    return ans
}