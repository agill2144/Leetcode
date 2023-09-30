/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    // we need to return most FREQUENT subtree sum
    // i.e the sum of subtree that we have seen the most amount of times
    // if all the sums are uniq, i.e the number of times all subtree sum is 1, we need to return all of those sums
    
    // keep track of {subtreeSum: countTimes} in a freqMap
    // and since we care about the subTreeSum seen the most ( i.e the count being the highest )
    // keep track of the highest counter
    sumCounter := map[int]int{}
    maxCount := 0
    
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        total := left+right+r.Val
        sumCounter[total]++
        if val := sumCounter[total]; val > maxCount {
            maxCount = val
        }
        return total
    }
    dfs(root)
    out := []int{}
    // collect all the sums that have maxCount
    for sum, count := range sumCounter {
        if count == maxCount {
            out = append(out, sum)
        }
    }
    return out
}