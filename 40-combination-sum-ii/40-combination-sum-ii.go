/*
    We are making subsequences, order of elements in a subsequence must match the order from input array
    We also cannot use the same element more than once
    We also have dupes
    
    approach: DFS + Backtracking ( for loop based recursion )
    - If we draw out the recursion tree with duped elements, you will see duplicate combinations being formed
    - So how do we not generate duplication combinations?
    - The issue is when we are adding items into our combination path, this item may be repeated more than once
    - Therefore leading to same exact combinations being generated more than once
        - Same exact, may be rearranged but same number of digits with same complements
    - So we need to make sure the items we add to our path ARE ALWAYS UNIQUE
    - How do we de-dupe?
        - hashset
            - Yes, although, we still need to use the dupe number $duped times, but just cannot branch out a path using the same duped numbers
            - so this wont work
        - map
            - Yes, almost
            - The issue with this is that hashmaps are not indexed
            - And for looping over a hash map does not provide any control over which key to start from
            - Why is this a problem?
            - We are generating subsequences!
            - Subsequences order matters, they are part of the array that may/may-not contain all elements
            - but they must FOLLOW THE ORDER OF INPUT ORDER
            - Hashmaps will not gurantee us this requirement
            - Therefore we will first convert to a freqMap
            - then transform the freqMap to a 2D array [ [$key, $freq] ]
            - Then we can run a normal for loop based recursion on this 2d array
    - So once we have a 2D, array, we will do the classic for loop based recursion
        - For loop based recursion - since we may not have to use all elements
    - What about the start pointer?
    - yes, we will have a start pointer starting from idx 0 of the 2D array
    - However loop in the dfs call will check if this idx has anymore $freq left
    - If not, continue
    - otherwise:
        - action: add this idx element to our path, decrement its $freq
        - recurse:
            - the new start pointer, will AGAIN BE the SAME i
            - I thought we are not allowed to use the same element more than once
            - We are not, the deduped array is a 2d freqMap array
            - We may still have more $freq left from this ith value, so therefore use the same i as new start position
        - backtrack: remove last element from our path and increment the $freq value of this idx to try another combination.
    Time: 
        o(2n) to build freqMap + deduped array
        +
        o(2^n) ? for building out all possible combinations
        
    Space:
        o(n) to build freqMap + deduped array + recursion stack
*/
func combinationSum2(candidates []int, target int) [][]int {
    if candidates == nil || len(candidates) == 0 || target == 0 {
        return nil
    }
    n := len(candidates)
    
    freqMap := map[int]int{}
    for i := 0; i < n; i++ { freqMap[candidates[i]]++ }
    
    deduped := [][]int{}
    for k, v := range freqMap { deduped = append(deduped, []int{k, v})}
    
    result := [][]int{}
    var dfs func(start int, path []int, t int)
    dfs = func(start int, path []int, t int) {
        // base
        if t <= 0 {
            if t == 0 {
                newL := make([]int, len(path))
                copy(newL, path)
                result = append(result, newL)
            }
            return
        }
        
        // logic
        for i := start; i < len(deduped); i++ {
            if deduped[i][1] == 0 {continue}
            // action
            path = append(path, deduped[i][0])
            deduped[i][1]--
            
            // recurse
            dfs(i, path, t-deduped[i][0])
            
            // backtrack
            deduped[i][1]++
            path = path[:len(path)-1]
        }
    }
    
    dfs(0, nil, target)
    return result
}



