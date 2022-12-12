func combinationSum(candidates []int, target int) [][]int {
    allPaths := [][]int{}
    helper(0, candidates, target, nil, &allPaths)
    return allPaths
}

func helper(start int, candidates []int, t int, path []int, allPaths *[][]int) {
    // base
    if t <= 0 {
        if t < 0 {return}
        newL := make([]int, len(path))
        copy(newL, path)
        *allPaths = append(*allPaths, newL)
        return
    }
    
    // logic
    for i := start; i < len(candidates); i++ {
        path = append(path, candidates[i])
        helper(i, candidates, t-candidates[i],path, allPaths)
        path = path[:len(path)-1]
    }
}