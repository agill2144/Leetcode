/*
    approach: brute force using extra space
    - Create a leftSum array
    - loop over nums arr from left to right
    - calc running left sum and add to leftSum array
        - prevLeftSum (i-1) + curr number nums[i]
    - Create a rightSum array
    - loop over nums arr from right to left
    - calc running right sum and add to rightSum array
        - prevRightSum (i+1)  + curr number nums[i]
    - wherever the 2 elements match in leftSum and rightSum, thats the pivot idx
    
    Time: o(2n), i.e; o(n)
    Space: o(2n), i.e; o(n)
*/


// func pivotIndex(nums []int) int {
//     leftSum := []int{nums[0]}
//     for i := 1; i < len(nums) ;i++ {
//         leftSum = append(leftSum, leftSum[i-1]+nums[i])
//     }
//     rightSum := make([]int, len(nums))
//     rightSum[len(rightSum)-1] = nums[len(nums)-1]
//     for i := len(nums)-2; i >=0 ; i-- {
//         rightSum[i] = nums[i] + rightSum[i+1]
//     }
    
//     for i := 0; i < len(nums); i++ {
//         if leftSum[i] == rightSum[i] {
//             return i
//         }
//     }
    
//     return -1
// }

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
func pivotIndex(nums []int) int {
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