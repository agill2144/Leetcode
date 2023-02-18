/*
    approach: bottom up dp 
    - space optimized if k is always going to < len(arr)
    time: o(n)
    space: o(k)
*/
func maxSumAfterPartitioning(arr []int, k int) int {
    dp := make([]int, k)
    dp[0] = arr[0]
    out := dp[0]
    for i := 1; i < len(arr); i++ {
        max := 0
        maxTotal := 0
        for j := i; j > i-k && j >= 0; j-- {
            if arr[j] > max {max = arr[j]}
            total := max * (i-j+1)
            if j > 0 {
                total += dp[(j-1)%k]
            }
            if total > maxTotal {maxTotal = total}
        }
        dp[i%k] = maxTotal
        if dp[i%k] > out {out=dp[i%k]}
    }
    return out
}


/*
    approach: bottom up dp
    - Solve smaller sub problem from bottom up
        - Decisions: k - for each ith element we have k choices, k partitions
        - We will look back from ith element and see what makes sense for this ith element 
        - The answer for each ith element is max sum in a partition and we could have k partitions for each ith element
        - Which ever partition creates maxSum, that will be the answer for each ith element.
        
    - What do we need to do for each subproblem?
        - Find max partition sum 
    - How are we going to find the max partition sum in bottom up dp?
        - For each ith element, we will walk back ( as long as we can )
        - And form each partition with this ith element
            - While making sure we are saving max and using that max to compute the sum.
        - How are we going to form k partitions with ith element?
            - Since bottom up dp means, looking back
            - we will start a j loop and walk back i-j steps
            - each iteration of this j loop is a partition of size 1 to k
            - in each iteration, we will save max seen so far in this window
            - then we will get the partition size we are solving for ( i-j+1 -- windowSize )
            - then we will multiple maxSeenSoFarInThisWindow * partitionSize to get the max sum possible with current parition
            - However we are not done yet..
            - We need to take the optimal answer for immediate left element outside of this partition ( left element just outside this partition )
            - Where we would we find the optimal answer for an element thats outside the partition window
            - THIS IS THE SUBPROBLEM THATS ALREADY SOLVED, so dp[j-1] will give you exactly that
            - Add previous subproblem answer to this paritionSum and save maxPartitionSum
        - Once we are outside of the j loop that runs k times for each ith element
        - We will save the maxPartitionSum formed by the jth loop for this ith element
        

        - Are we re-using previous/left element decisions ?
            - Yes, since we are doing bottom up,
            - Then for each ith element, we have k choices
            - In this bottom approach for each ith element, we will look left/previous element
            - Then we will make k partitions for this ith element
            - Whatever is left outside of parition, we will add that in the current partition calculation.
            - So we will have an ith element
            - and lets say k = 2
            - Then ith element can have 1 to 2 parititons
            - Then we will loop from i and back as long as j > i-k and j >= 0
            - While we are making each partition, we will keep track of the max seen so far while looping back
            - In each iteration, we will take current max seen so far * paritionSize ( i-j+1 ) - i.e windowSize
            - Then whatever elements are to the left ( outside ) of this partition, they already have an answer in our dp array
            - So will pick whatever was the max answer for elements outside of partition Array (dp[j-1]) 
                - if j-1 < 0, then we have to previous subproblem to re-use their answer for. move on with whatever is in hand
                - THIS IS HOW WE WILL USE PREVIOUSLY SOLVED SUBPROBLEMS TO ANSWER CURRENT SUBPROBLEM
        - Finally the max parition sum of size k will be saved in the last idx on dp array.
*/
// time: o(nk)
// space: o(n)
// func maxSumAfterPartitioning(arr []int, k int) int {
//     n := len(arr)
//     dp := make([]int, n)
//     dp[0] = arr[0]
    
//     for i := 1; i < n; i++ {
        
//         maxInPart := 0
//         for j := i; j > i-k && j >= 0 ; j-- {
//             if arr[j] > maxInPart { maxInPart = arr[j] }
//             partSize := i-j+1
//             partitionSum := maxInPart * partSize
//             if j-1 >= 0 {
//                 partitionSum += dp[j-1]
//             }
//             if partitionSum > dp[i] {dp[i] = partitionSum}
//         }
//     }
    
//     return dp[len(dp)-1]
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }
