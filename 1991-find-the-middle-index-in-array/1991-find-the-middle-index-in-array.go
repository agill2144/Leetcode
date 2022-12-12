/*
    How is this a prefix/running sum?
    - we are looking for a idx whose sum on left ( excluding the idx ) is equal to the sum of elements on right side of this idx
    - if we had total sum and then for each idx, we can calc the leftSum ( i.e running sum ) and we can get the right sum by
    - totalSum - leftSum -- this gives the right sum ( given we have total and leftSum )
    - but we have to exclude the idx under evaluation to be the pivot idx
    - therefore to exclude the pivot idx, we will do totalSum - leftSumSoFar - currentValAtIdx
    - once we do that , we have leftSum and we have rightSum 
    - if both are equal, this means our current idx is separator/pivot idx

    time: o(2n) - o(n)
    space: o(1)
*/
func findMiddleIndex(nums []int) int {
    totalSum := 0
    for i := 0; i < len(nums); i++ {
        totalSum += nums[i]
    }
    
    leftSum := 0
    for i := 0; i < len(nums); i++ {
        rightSumExcludingI := totalSum-leftSum-nums[i]
        if rightSumExcludingI == leftSum {
            return i
        }
        leftSum += nums[i]
    }
    return -1
}