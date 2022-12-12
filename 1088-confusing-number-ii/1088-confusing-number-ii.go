/*

    approach: DFS
    - First, there are only handful numbers from 0 to 9 that can be flipped into valid numbers
        - 0:0,1:1,6:9,8:8,9:6
    - We are given N, so instead of going from 0 till N , 1 number at a time, we dont have to do that
    - Because the numbers that cannot be flipped at all, there is no need to evaluate those numbers right?
    - This means, the numbers we need to form are going to be with digits that can be flipped into valid numbers.
    - How do we check if a number formed is a confusing number?
        - We have to reverse the number while flipping each digit to a new number.
        - For example, 16
        - take the last digit ( using %10 ) - we get 6
        - if we flip 6 - we get 9 - so valid
        - then flip 1, we get 1 
        - concatanate the revers flips of each digits of 16 and we get 91
        - 91 != 16 - therefore yes, its a confusing number
        - 11 != 11 - false, 11 cannot be a confusing number because after flipping we get the same original number back
    - But how do we form all the numbers till N while only making combinations with the valid digits that can be flipped?
    - combinations == for loop based recursion with backtracking 
        - Do we have any reference data structure to backtrack tho?
        - when we want our recursion to go back to a previous parent, we want to go back to the previous number formed
        - which is an integer and ints are not reference based, therefore no explicit backtracking needed, recursion will handle it.
        - NOTE: just make sure to not modify the the number in recursion stack or else when we go back to that parent, we wont have the original number
            - in this case if we do mutate, then we do have to backtrack our mutation action
            - therefore just create a new var for concatnating any new number into current passed number from args.
        - Every time we form a number, we can evaluate whether its confusing number
        - So we will recursively call our dfs func from within our for loop
        - We will first check if this number is > N - if yes, return back to parent recursion call
        - Then we will evaluate whether the current passed number is a confusing number or not
            - if yes, count++
            - Potentianl follow up ans: What if we are asked to also return those numbers? - this is where we would be saving them.
        - And then run a for loop on the map again to add more numbers and evaluate those combinations.
    
    Time: o(n) 
    - we create n/2 combinations ( since our flipMap does not contain all numbers from 0 to 9)
    Space: recursion stack - so worse case o(k) where k is the number of digits in input n 
        - since we never let a recursion call who is greater than N itself - technically k+1
        - the flipMap is constant

*/
func confusingNumberII(n int) int {
    
    count := 0
    flipMap := map[int]int{
        0:0,1:1,6:9,8:8,9:6,
    }
    var dfs func(curr int)
    dfs = func(curr int) {
        // base
        if curr > n {return}
        
        // logic
        if isConfusingNum(curr, flipMap) {
            count++
        }
        for k, _ := range flipMap {
            // action - append new int to curr
            // curr = curr * 10 + k
            // recurse
            // if curr != 0 {
            //     dfs(curr)
            // }
            // backtrack - remove last digit that was initially added at the start of the loop
            // curr /= 10
            
            // the above is not needed if we do not modify curr
            newCurr := curr * 10 + k
            if newCurr != 0 {
                dfs(newCurr)
            }
            
        }
    }
    dfs(0)
    return count
}

func isConfusingNum(n int, flipMap map[int]int) bool {
    orig := n
    result := 0
    for n != 0 {
        last := n % 10
        n = n / 10
        lastFlipped := flipMap[last]
        result = result * 10 + lastFlipped
    }
    return orig != result
} 