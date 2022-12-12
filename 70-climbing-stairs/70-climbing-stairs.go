/*
    approach: Bottom up DP
    - We have to explore all paths, and somewhere we will see repeated subproblems
    - This is like unique paths dp question
    - We will treat each stair as our final destination and solve for this subproblem
    - What are we solving for: We are determining number of ways to get to this step
    - And the number of ways have constraints, you can come back from 2 steps back or 1 step back
        - because we can either climb 1 step or 2 steps
        - so if we are solving for a step i, then we need to add number of ways previous step can be reached and number of ways 2 step back steps can be reached
        - why? BECAUSE TO GET TO THIS STEP, WE CAN ONLY COME BACK FROM 1 STEP BACK OR WE CAN COME FROM 2 STEPS BACK
        - ok this makes sense for some ith element thats in the middle..
        - how would we solve this for 1st stair...
            - The 1st stair only have 1 viable option, that is 1 step - so therefore total number of ways to get to step1 WILL ALWAYS BE 1
            - then we go to step 2
            - step2 can be reached if we took 1 step from its previous step + 2 step from its 2 step back stair( which does not exist so only 1 option )
            - then we solve for step3
            - how many ways can step3 be reached from the left side 
                - Remember we can either climb 1 step or 2 steps
                - Which means, step3 can be reached in 1 step from stair2
                - Which means, step3 can be reached in 2 steps from stair1
                - So total ways to reach to step3 would the sum of above 2 values
    
    Time: o(n)
    Space: o(n) - since we will use and maintain a dp array
*/
// time : o(n)
// space : o(n)
// func climbStairs(n int) int {
    
//     dp := make([]int, n+1)
//     dp[0] = 1
//     dp[1] = 1
    
//     for i := 2; i < len(dp); i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }
    
//     return dp[len(dp)-1]
    
// }



// we only need previous values
// how would we store the new total and where?
// well it turns out that the new total is the new twoStepsBack value
// but first promote the twoStepBack value to 1 step back position
// and then store the new total in twoStepBack position
// This way our space is constant
// time : o(n)
// space : o(1)
func climbStairs(n int) int {
    
    dp := []int{1,1}
    
    for i := 2; i <= n; i++ {
        oneStepBack := dp[0]
        twoStepBack := dp[1]
        dp[0] = twoStepBack
        dp[1] = oneStepBack+twoStepBack
    }
    
    return dp[len(dp)-1]
    
}