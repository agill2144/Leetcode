/*
    approach: for loop based recursion / backtracking
    - like combination sum
    - form all possible combinations
    - at each recursive call, we have an ans to save
    
    time ; n* 2^n
    n * because at each recursive call, we create another array , deepCp and save the ans
    space ;
    o(n) recursion stack * o(n) for new deepCp array ; n^2 
*/
func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, curr := range out {
            newP := make([]int, len(curr))
            copy(newP, curr)
            newP = append(newP, nums[i])
            out = append(out, newP)
        }
    }
    return out
}