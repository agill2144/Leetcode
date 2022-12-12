func subsetsWithDup(nums []int) [][]int {
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        val, ok := freqMap[num]
        if !ok {
            freqMap[num] = 1
        } else {
            freqMap[num] = val+1
        }
    }
    
    deduped := [][]int{}
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        val, ok := freqMap[num]
        if ok {
            deduped = append(deduped,[]int{num,val})
            delete(freqMap,num)
        }
    }
    
    result := [][]int{}
    var dfs func(idx int, path []int)
    dfs = func(idx int, path []int) {
        // base 
        
        
        // logic
        newL := make([]int, len(path))
        copy(newL, path)
        result = append(result, newL)
        
        for i := idx; i < len(deduped); i++ {
            elePair := deduped[i]
            ele := elePair[0]
            freq := elePair[1]
            if freq == 0 {continue}
            
            // action
            path = append(path,ele)
            elePair[1]--
            
            // recurse
            dfs(i,path)
            
            // backtrack
            path = path[:len(path)-1]
            elePair[1]++
            
        }
    }
    dfs(0,[]int{})
    return result
}