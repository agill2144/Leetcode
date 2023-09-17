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
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        newPath := make([]int, len(path))
        copy(newPath, path)
        out = append(out, newPath)
        
        
        // logic
        for i := start; i < len(nums); i++ {
            // action
            path = append(path, nums[i])
            dfs(i+1, path)
            path = path[:len(path)-1]
        }
    }
    dfs(0, nil)
    return out
}