/*
    approach: dfs + backtracking
    - This is a variation of permutations but the input array may contain duplicate elements
    - So why cant we use the same for-loop-based recursion + backtracking on this problem?
        - How did we solve the first permutation problem using backtracking?
        - We had a for loop based recursion in our dfs
        - We deployed a dfs on each ith element.
        - We also had a start pointer telling where to start the for loop from
            - Why ?
            - if we start our for loop from 0th idx in all recursion calls, we will have all the same numbers again
            - once we have selected a number, we need to move to the next number and explore that number depth first.
        - Then to create all possible permutations, in each for loop iteration, we swapped ith element with start pointer
        - This way we explored placing each element at each position - for all elements
        - The reason why we cannot use the above directly on this question is because we have dupes in the input array.
        - The dupes lead to SAME EXACT PERMUTATIONS BEING GENERATED multiple times
        - Draw out the for loop recursion tree to see this in action
    - How to handle the dupe permutations being generated
        - The dupes are being generated because we have dupe elements in the input array
    - Instead we will transform the array into a freqMap
    - Then our for loop based recursion will run on the freqMap
    - We wont need a start pointer because maps are not indexed.
    - Our for loop will simply check whether this keys value/count > 0
    - ACTION : If it is, we have a viable number to use, so add it to our path and recurse and reduce its frequency by 1
    - Then we will recurse
        - Our base case will eval whether we have a possible uniq permutation yet or not
        - We will know this if the size of our path array == len of input array
    - Then once the child call finishes, and goes back to our parent, since this is dfs, do we need to backtrack anything?
        - Do we have any reference data structures being changed in each iteration?
        - Yes, we have 2. One is the path array and other is the freqMap counter being changed for each key
    - Backtrack: remove the last element from path and increment the current k freq count by 1
    
    - How did freqMap help in ensuring we do not generate duplicate permutations
        - If the input array does not have dupes, the regular for loop based recursion + backtrack will generate all possible permutations.
        - Since the input array has dupes, we made same decisions on same duplicated elements/items
        - By putting these elements into a freqMap, the keys were kept unique, hence were back to the same original problem - generate permutations with unique elements
        - with the only difference being is that our for loop would run on the map instead of the array.

    Time: o(2^n)
    Space: o(n)
*/
func permuteUnique(nums []int) [][]int {
    
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freqMap[nums[i]]++
    }
    
    result := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int) {
        // base
        if len(path) == len(nums) {
            newL := make([]int, len(path))
            copy(newL, path)
            result = append(result, newL)
            return 
        }
        
        // logic
        for k, v := range freqMap {
            if v == 0 {continue}
            // action
            path = append(path, k)
            freqMap[k]--
            
            // recurse
            dfs(path)
            
            // backtrack
            path = path[:len(path)-1]
            freqMap[k]++
        }
    }
    
    dfs(nil)
    return result
    
}