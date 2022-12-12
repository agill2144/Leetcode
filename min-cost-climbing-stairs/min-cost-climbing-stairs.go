/*
    We have to find the min cost to climb all the stairs
    - Stairs in this example is an array
    - Bottom of the stair is idx 0 and top of the stair is past the last element
    - We can either take 1 step or 2 steps
        - NOTICE : WE ARE GIVEN CHOICES - DP hint
    - And use a step in of the stairs has a cost
    - And we have to strategically climb all the stairs while minimizing cost to go all the way up
        - NOTICE : Asking "minimum" -- something optimal - another DP hint
    - Which likely means there are many ways to go about climbing these stairs
    - However we also have 1 extra choice
        - We can either start from the bottom of the stairs ( idx 0 )
        - Or we can start from idx 1 ( i.e we do not have to count in the cost of idx 0 step )
        - whichever costs less
    
    - After drawing the decision tree, you will notice repeated subproblems
    - Which is a clear indication of dp
    - We can start by solving the smallest subproblem from top down
    - Which means, we can start from the top of the stairs, and work back down
    - At each stair step, we have 2 choices
        - take 1 step
        - take 2 steps
    - Which ever one results in smaller cost, we will pick that ofc - since we are trying to minimize the cost
    - So for each element from the back of the costs array ( 2nd last element , since last element cost would not be different if we take 1 or 2 step )
        - We will check whether
            - is it cheaper to take 1 step ( currentCost + nextStepCost )
            - is it cheaper to take 2 step ( currentCost + 2StepsNextCost ) -- if 2 steps goes out of bounds, we can assume 0 for 2 steps price
        - Which ever one is cheaper we will set that price for that step
    - And work back to the front of the array
    - When we get to the first step ( 0th idx ) - we have more choices here to make
            - is it cheaper to take 1 step ( currentCost + nextStepCost )
            - is it cheaper to take 2 step ( currentCost + 2StepsNextCost ) -- if 2 steps goes out of bounds, we can assume 0 for 2 steps price
            - or is cheaper to skip this step: min(nextStepCost, min(nextStepCost,2NextStepCost)+currentCost )
    
    - Finally at the end we have bubbled our answer up in the 0th idx
    - We took the cheapest decision on each step by evaluation whether its better to take 1 or 2 steps - from the back of the array 
    - We essentially solved the smallest subproblem for each step - and worked up to bubble up the final ans.
    
    time: o(n) 
        - n is the len of cost array
    space: o(n)
        - n is the len of cost array
        - we allocated extra dp array 
        - could be optimized to o(1) space, if we can mutate the cost array
*/
// func minCostClimbingStairs(cost []int) int {
//     dp := make([]int, len(cost))
//     dp[len(dp)-1] = cost[len(cost)-1]
    
//     for i := len(cost)-2; i >= 0; i-- {
//         oneStepCost := dp[i+1]
//         twoStepCost := 0
//         if i+2 < len(cost) {
//             twoStepCost = dp[i+2]
//         }
//         currentCost := cost[i]
        
//         if i == 0 {
//             ifWeChoose0thIdx := min(oneStepCost,twoStepCost)+currentCost
//             dp[i] =  min(oneStepCost,ifWeChoose0thIdx )
//         } else {
//             dp[i] = int(math.Min(float64(oneStepCost), float64(twoStepCost)))+currentCost
//         }
//     }
//     return dp[0]
// }

/*
    approach: bottom up dp
    - solve the smallest subproblem and use previous decisions to enhance smallest subproblem ans
*/
func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost))
    dp[0] = cost[0]
    
    for i := 1; i < len(dp); i++ {
        currStepCost := cost[i]
        oneStepBackCost := dp[i-1]
        twoStepBackCost := 0
        if i-2 >= 0 {
            twoStepBackCost = dp[i-2]
        }
        dp[i] = min(currStepCost+oneStepBackCost, currStepCost+twoStepBackCost)
    }
    return min(dp[len(dp)-1], dp[len(dp)-2])
}


func min(x, y int) int{
    if x < y {return x}
    return y
}