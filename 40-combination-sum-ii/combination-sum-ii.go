// tc = o(n) + o(n) + 2^n
// sc = o(2n) + o(t)
// o(t) = recursion stack stops after runningSum exceeds target, 
//          assume all values are 1s, 
//          therefore reucrsion stack will be of size $target)
func combinationSum2(candidates []int, target int) [][]int {
    freqM := map[int]int{}
    for i := 0; i < len(candidates); i++ {freqM[candidates[i]]++}
    freq := [][]int{}
    for k, v := range freqM {freq = append(freq, []int{k,v})}
    out := [][]int{}
    var dfs func(start , sum int, path []int) 
    dfs = func(start, sum int, path []int) {
        // base
        if sum >= target {
            if sum == target {
                newP := make([]int, len(path))
                copy(newP, path)
                out = append(out, newP)
            }
            return
        }

        // logic
        for i := start; i < len(freq); i++ {
            if freq[i][1] > 0 {
                // action
                sum += freq[i][0]
                freq[i][1]--
                path = append(path, freq[i][0])
                // recurse
                dfs(i, sum, path)
                // backtrack
                path = path[:len(path)-1]
                freq[i][1]++
                sum -= freq[i][0]
            }
        }
    }
    dfs(0, 0, nil)
    return out
}