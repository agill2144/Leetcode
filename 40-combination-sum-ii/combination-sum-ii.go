func combinationSum2(nums []int, target int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    var dfs func(start, sum int, path []int)
    dfs = func(start, sum int, path []int){
        // base
        if sum == target {
            newL := make([]int, len(path))
            copy(newL, path)
            out = append(out, newL)
            return            
        }
        if start == len(nums) {return}

        // logic
        for i := start; i < len(nums); i++ {
            if i > start && nums[i] == nums[i-1] {continue}
            if sum + nums[i] > target {return}
            // action
            sum += nums[i]
            path = append(path, nums[i])
            // recurse
            dfs(i+1, sum, path)
            // backtrack
            path = path[:len(path)-1]
            sum -= nums[i]
        }
    }
    dfs(0,0, []int{})
    return out
}
// func combinationSum2(nums []int, target int) [][]int {
//     m := map[int]int{}
//     n := len(nums)
//     for i := 0; i < n; i++ {m[nums[i]]++}
//     freq := [][]int{} // [ [val, count] ]
//     for k, v := range m {
//         freq = append(freq, []int{k, v})
//     }
//     out := [][]int{}
//     var dfs func(start int, path []int, sum int)
//     dfs = func(start int, path []int, sum int) {
//         // base
//         if sum == target {
//             newL := make([]int, len(path))
//             copy(newL, path)
//             out = append(out, newL)
//             return
//         }
//         if sum > target || start == len(freq) {return}

//         // logic
//         for i := start; i < len(freq); i++ {
//             if freq[i][1] == 0 {continue}
//             // action
//             sum += freq[i][0]
//             freq[i][1]--
//             path = append(path, freq[i][0])
//             // recurse
//             dfs(i, path, sum)
//             // backtrack
//             path = path[:len(path)-1]
//             freq[i][1]++
//             sum -= freq[i][0]
//         }
//     }
//     dfs(0, nil, 0)
//     return out
// }