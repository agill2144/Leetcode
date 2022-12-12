/*
    - We have to minimize cost for dates we are traveling
        - Note: "minimize cost" - asking for something optimal ( maybe DP )
    - We also have choices to pick from for each day
        - 1 day, 7 days, and 30 days pass with different cost for each.
        - Note: "choices" - maybe dp
    - We are given a days array which contains asc order of dates we will travel on
    - These dates may not be consecutive in value, there maybe gaps like [1,2,5,12,20]
    - despite the gaps, they are in ascending order
    - costs array is given to us [2,7,30] <2DayPass>,<7DayPass>,<30DayPass>
    - For each day we have to make a decision and minimize our overall cost
    - After drawing the decision tree, you will see repeated subproblems
    - Therefore we can explore DP
    - Solve the smallest subproblem and bubble your answer in a dp array/matrix
    
    
    approach: top down DP from back of the array
    - We can create a dp array of size $maxVal+1 in days array ( i.e the last element )
        - why maxVal ?
            - Since the days array have gaps, it will be harder to determine what cost to pick based on current date and the choice
            - For example, if the days array is: [1,2,5,12,20]
            - We will create a dp array of size 20+1 = 21
            - [0,0,0,0,0,0,0,0,0, 0,  0,  0,  0,  0,  0,  0 , 0,  0,  0,  0]
               1,2,3,4,5,6,7,8,9,10, 11,  12, 13, 14, 15, 16, 17, 18, 19, 20
            - So that each idx represents a potential date from the days array
            - This fills up the gap thats input days array may have
            - This makes it easier calculate whats cheapest for a subproblem
    - For a bare minimum subproblem ( from right of the dp array and back - top down )
    - For each subproblem we have 3 choices
        1. 1 day pass
        2. 7 day pass
        3. 30 day pass
    - If we are solving top down ( looking right ) for an ith position
        - if we pick 1 day pass - than total cost would be 2(1dayPass) + totalCostSoFar(i+1)
        - if we pick 7 day pass - than total cost would be 7(7dayPass) + totalCostWhileNotIncludingTheNext7Days
            - Why not include the next 7 days?
            - Why should we?
                - if we are determining the total cost from ith day, than if we buy a 7 day pass
                - we need to see until what date this pass will cover and if there are remaning days, whats their cost
                - dp[i+7] if inbound
        - if we pick 30 day pass - than totalCost would be 15(30DayPass) + totalCostAfterExcludingTheNext30DaysCoveredByThisPass
    - This way we will start from the back of the array ( 2nd last element )
        - the last element will by default pick the cheapest price and it does not have next days to do any other math with
    - And work back to the start of the array and bubble up our ans to 0th idx\
        
        TLDR - top down from the back of array solving each subproblem ( min of 3 choices for each subproblem )
    
    time: o(n)
    space: o(n+1)  
        i.e:  o(n)
*/
// top down from the back of the days array - bubbling the ans to idx 0
// func mincostTickets(days []int, costs []int) int {
    
//     travelDates := map[int]struct{}{}
//     for i := 0; i < len(days); i++ {
//         travelDates[days[i]] = struct{}{}
//     }
    
//     // bottom up dp
//     // each subproblem has 3 choices, make the best one for each ( i.e min cost )
//     dp := make([]int, days[len(days)-1])
//     n := len(dp)
//     oneDayPass := costs[0]
//     sevenDayPass := costs[1]
//     thirtyDayPass := costs[2]
//     dp[len(dp)-1] = min(oneDayPass,min(sevenDayPass,thirtyDayPass))

//     for i := len(dp)-2; i >= 0; i-- {
//         travelDate := i+1
//         _, traveling := travelDates[travelDate]
//         if !traveling { dp[i] = dp[i+1]; continue }
    
//         oneDayTotalCost := oneDayPass + dp[i+1]
//         sevenDayTotalCost := sevenDayPass
//         if i+7 < n { sevenDayTotalCost = sevenDayPass + dp[i+7] }
//         thirtyDayTotalCost := thirtyDayPass
//         if i+30 < n {thirtyDayTotalCost = thirtyDayPass + dp[i+30] }
//         dp[i] = min(oneDayTotalCost,min(sevenDayTotalCost,thirtyDayTotalCost))
//     }
//     return dp[0]
// }

/*
    approach: bottom up
    - We can create a dp array of size $maxVal+1 in days array ( i.e the last element )
        - why maxVal ?
            - Since the days array have gaps, it will be harder to determine what cost to pick based on current date and the choice
            - For example, if the days array is: [1,2,5,12,20]
            - We will create a dp array of size 20+1 = 21
            - [0,0,0,0,0,0,0,0,0, 0,  0,  0,  0,  0,  0,  0 , 0,  0,  0,  0]
               1,2,3,4,5,6,7,8,9,10, 11,  12, 13, 14, 15, 16, 17, 18, 19, 20
            - So that each idx represents a potential date from the days array
            - This fills up the gap thats input days array may have
            - This makes it easier calculate whats cheapest for a subproblem
    - To solve a bare subproblem ( from left of the dp array - bottom up - compare/look-at previous decisions to make a better decision )
    - For each subproblem we have 3 choices
        1. 1 day pass
        2. 7 day pass
        3. 30 day pass
    - If we are solving bottom up ( looking left ) for an ith position
        - if we pick 1 day pass - than total cost would be 2(1dayPass) + totalCostSoFar(i-1)
        - if we pick 7 day pass - than total cost would be 7(7dayPass) + totalCostAfterExcludingThePrev7DaysCoveredByThisPass
            - Why not include the previous 7 days?
            - Why should we?
                - if we are determining the total cost from ith day, than if we buy a 7 day pass
                - we need to see from date we should have bought this and if there are more prev remaning days, whats their cost
                - dp[i-7] if inbound
        - if we pick 30 day pass - than totalCost would be 15(30DayPass) + totalCostAfterExcludingThePrev30DaysCoveredByThisPass
    - This way we will start from the front of the array ( 2nd element )
        - for the first date ( 0th idx val in days array == idx in dp array ) element will by default pick the cheapest price and it does not have prev days to do any other math with
    - And work forward to the end of the array and bubble up our ans to last idx
        
        TLDR - bottom up from the front of array solving each subproblem while looking back (left) for better decision making ( min of 3 choices for each subproblem )
    
    time: o(n)
    space: o(n+1)  
        i.e:  o(n)
*/

func mincostTickets(days []int, costs []int) int {
    dp := make([]int, days[len(days)-1]+1)
    
    oneDayPass := costs[0]; sevenDayPass := costs[1]; thirtyDayPass := costs[2]
    dp[days[0]] = min(oneDayPass, min(sevenDayPass, thirtyDayPass))
    
    dates := map[int]struct{}{}
    for i := 0; i < len(days); i++ {
        dates[days[i]] = struct{}{}
    }
    
    
    for i := days[0]+1; i < len(dp); i++ {
        _, ok := dates[i]
        if !ok {
            dp[i] = dp[i-1]
            continue
        }
        
        oneTotal := oneDayPass + dp[i-1]
        sevenTotal := sevenDayPass
        if i-7 >= 0 {
            sevenTotal +=  dp[i-7]
        }
        thirtyTotal := thirtyDayPass
        if i-30 >= 0 {
            thirtyTotal += dp[i-30]
        }
        
        dp[i] = min(oneTotal, min(thirtyTotal,sevenTotal))
    }
    return dp[len(dp)-1]
}

func min(x, y int) int{
    if x < y {return x}
    return y
}