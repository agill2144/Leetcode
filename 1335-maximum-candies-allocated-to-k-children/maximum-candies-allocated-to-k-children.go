func maximumCandies(candies []int, k int64) int {
    // koko flashbacks 
    totalCandies := 0
    left := 1
    right := math.MinInt64
    for i := 0; i < len(candies); i++ {
        totalCandies += candies[i]
        right = max(right, candies[i])
    }
    if int64(totalCandies) < k {return 0}
    ans := 0
    for left <= right {

        // if we are allocation mid number of candies per student
        mid := left + (right-left)/2

        // how many piles can we create in total ?
            // how many piles can we create from each candies[i]th pile ?
            // if there are 5 candies in a pile
            // our mid is also 5
            // we can create 1 pile 
            // if there are 10 candies in a pile
            // our mid is 5
            // we can create/divide this pile into 2 piles of 5 each
            // 10/5 = 2            

        // why are we looking for number of total piles ?
        // - "You should allocate piles of candies to k children such that each child gets the same number of candies"
        // therefore number of piles == number of students 

        var count int64 = 0
        for i := 0; i < len(candies); i++ {
            count += int64(candies[i])/int64(mid)
        }

        // when does mid work?
        // when we were able to allocate mid number of candies for atleast k students
        if count >= k {
            ans = mid
            // but we want maximum such ans
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}