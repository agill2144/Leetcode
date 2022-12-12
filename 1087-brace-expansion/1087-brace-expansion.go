/*
    approach: backtrack
    - what are we backtracking?
    - the path we form from input string
    - We will first convert our input string to be parsable list of strings
    - This way we wont have to complicate our logic with what to when we run into different types of characters
    - Our main logic will simply work with a for loop
    - We will have a 2D nested array of characters
    - So our backtrack func will be hold the pointer to outter element
    - While inside each for backtrack call, a for loop will loop over each char within a nested array
    
    Time: exponential
    space: o(n) to expand input string into a nested array
*/  
func expand(s string) []string {
    blocks := [][]string{}
    i := 0
    // time:  o(n) to loop the string
    for i < len(s) {
        block := []string{}
        if s[i] == '{' {
            i++
            for s[i] != '}' {
                if s[i] == ',' {i++; continue}
                block = append(block, string(s[i]))
                i++
            }
        } else {
            block = append(block, string(s[i]))
        }
        i++
        sort.Strings(block)
        blocks = append(blocks, block)
        
    }
    result := []string{}
    var backtrack func(path string, start int)
    backtrack = func(path string, start int) {
        // base
        if start == len(blocks) {
            result = append(result, path)
            return
        }
        
        // logic
        for i := 0; i < len(blocks[start]); i++ {
            path += blocks[start][i]
            backtrack(path, start+1)
            path = path[:len(path)-1]
        }
    }
    backtrack("", 0)
    return result
}