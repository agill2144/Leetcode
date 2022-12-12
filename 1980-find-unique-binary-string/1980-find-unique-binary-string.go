/*
    approach: explore all paths (for loop based recursion with backtracking)
    - We will use a freqMap that contains 0 and 1 as our keys and values would be n where n is the len of nums array
    - n because all binary strings within array are also of len n
    - and we also have to look for a uniq binary string of size n
    - if n is the len of binary string we have to return
    - it means that we have n 1's and n 0's
    - if we used an array, than we will end up generating duplicate permutations - which would work too because we are checking if this permutation of binary string
    - exists in nums array or not
    - Which also means that we should convert our nums array to something that faster to search = so a map.
    - We will use backtracking to form every single permutation of the binary strings of len n and once we have a path that does not exist in nums map, return that path
    - otherwise keep being exhaustive and explore all permutations
    
    time: o(2^n)
    space: o(n)
*/
func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    
    set := map[string]bool{}
    for i := 0; i < n; i++ {
        set[nums[i]] = true   
    }
    m := map[string]int{"1":n, "0":n}
    
    var dfs func(path string) string
    dfs = func(path string) string {
        // base
        if len(path) == n {
            if _, ok := set[path]; !ok {
                return path
            }
            return ""
        }
        
        // logic
        for k, v := range m {
            if v == 0 {continue}
            path += k
            m[k]--
            if found := dfs(path); found != "" {
                return found
            }
            path = path[:len(path)-1]
            m[k]++
        }
        return ""
    }
    return dfs("") 
}