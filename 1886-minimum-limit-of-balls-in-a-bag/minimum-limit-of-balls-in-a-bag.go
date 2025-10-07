/*
    maximum number of balls in a bag. You want to minimize your penalty
    - atMax = mid
    - search for smallest such answer 
        - i.e if mid is a valid answer
        - search left, for a even smaller answer
    
    - atMax will be max num of balls allowerd per bag
    - to count num of operations needs to split a bag of nums[i] balls
        - ceil(nums[i] / atMax)-1
    - nums[i] = 8
    - atMax(or mid) = 2
    - num of operations needed to satisfy 2 balls per bag is ; (8/2)-1
    - its possible nums[i] = 7
    - hence it would take 1 more operation (not 1 less), hence the ceil 

    - nums[i] = 12
    - atMax = 2
    - if we think about splitting evenly, this would actually take 7 operations
    - but our formula suggests: (12/2)-1 = 5 operations
    - its NOT ABOUT SPLITTING EVENLY, ITS ABOUT SPLITTING OPTIMALLY

*/
func minimumSize(nums []int, maxOperations int) int {
    left := 1
    right := slices.Max(nums)
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        ops := 0
        for i := 0; i < len(nums); i++ {
            ops += int(math.Ceil(float64(nums[i])/float64(mid))-1.0)
        }
        if ops <= maxOperations {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}