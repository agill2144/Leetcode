/*
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
func permute(nums []int) [][]int {
    out := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int) {
        // base
        if len(path) == len(nums) {
            newP := make([]int, len(path))
            copy(newP, path)
            out = append(out, newP)
            return
        }
        
        // logic
        for i := 0; i < len(nums); i++ {
            if !listContains(path, nums[i]) {
                path = append(path, nums[i])
                dfs(path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(nil)
    return out
}

func listContains(list []int, ele int) bool {
    for _, x := range list{ 
        if x == ele {return true}
    }
    return false
} 