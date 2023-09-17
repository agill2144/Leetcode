/*
    because we need to place an element at each idx position
    [1,2,3]
    next permutation can be
    [1,3,2]
    then next
    [2,1,3]
    
    We can do this in-place
     s
    [1,2,3]
     i
    swap(i,s), recurse to next element (s+1) until we have reached the end
       s
    [1,2,3]
       i
         s
    [1,2,3]
         i
    now this is a permutation to save
    then recursion goes back to prev call and i moves forward
       s
    [1,2,3]
         i
        
    now inorder to generate [1,3,2], we have i and s at a perfect position, therefore swap(i,s)
    and start recursion from s+1
         s
    [1,3,2]
         i
    save and go back and backtrack the swap, i moves forward and goes out of bounds
       s
    [1,2,3]
           i
    
    Now we need to put 2 at idx 0
     s
    [1,2,3]
       i    
    the root recursive call will move i forward and we can again swap
    
    This is how we can swap in-place and make all possible permutations
    
    time = o(n*n!)
    space = o(n^2)
*/
func permute(nums []int) [][]int {
    out := [][]int{}
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(nums) {
            deepCp := make([]int, len(nums))
            copy(deepCp, nums)
            out = append(out, deepCp)
            return
        }
        
        // logic
        for i := start; i < len(nums); i++ {
            nums[i], nums[start] = nums[start], nums[i]
            dfs(start+1)
            nums[i], nums[start] = nums[start], nums[i]
        }
    }
    dfs(0)
    return out
}


/*
    In combinations, we place an element and started recursion for next element
    - in permutation, we place an element and start recursion AGAIN from start of the list
        and check whether we can add this element in path or not 
    - this wont work when there are duplicates of the same element because the check will say that an element already exists
    
    approach: backtracking using extra array
    - recursively place an element at each idx
    - we have to place each element at a specific index
    - place 1 at idx 0, then 2 at idx 0, then 3 at idx 0
    - and then repeat for idx 1 and remaining indicies
    - However we cannot place the same element twice right
    - therefore before placing an element at an idx, check if the arr already contains the same element
        - dont place it it element already exists, place only if it does not exist
    - once we have a path/arr of same size as input array, we have an answer to save
        - make a deepCopy of the path/array and save it

    time = o(n*n!)
    space = o(n^2) ( recursive stack and a deepcopy at the bottom most recursive call )
*/
// func permute(nums []int) [][]int {
//     out := [][]int{}
//     var dfs func(path []int)
//     dfs = func(path []int) {
//         // base
//         if len(path) == len(nums) {
//             newP := make([]int, len(path))
//             copy(newP, path)
//             out = append(out, newP)
//             return
//         }
        
//         // logic
//         for i := 0; i < len(nums); i++ {
//             if !listContains(path, nums[i]) {
//                 path = append(path, nums[i])
//                 dfs(path)
//                 path = path[:len(path)-1]
//             }
//         }
//     }
//     dfs(nil)
//     return out
// }

// func listContains(list []int, ele int) bool {
//     for _, x := range list{ 
//         if x == ele {return true}
//     }
//     return false
// } 