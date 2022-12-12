/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
/*

    approach: 2 Pass DFS
    1. 1 DFS pass to get the max depth
    2. 2 DFS pass to do the math using maxDepth from 1st Pass.
    
    Time: o(n)
    space: o(n)
*/
func depthSumInverse(nestedList []*NestedInteger) int {
    maxDepth := 0
    var maxDepthDfs func(nl []*NestedInteger, depth int)
    maxDepthDfs = func(nl []*NestedInteger, depth int) {
        // base
        if nl == nil {
            return
        }
        // logic
        maxDepth = int(math.Max(float64(maxDepth), float64(depth)))
        for _, n := range nl {
            if !n.IsInteger() {
                maxDepthDfs(n.GetList(), depth+1)
            }
        }
        
    }
    maxDepthDfs(nestedList, 1)
    
    total := 0
    var dfs func(nl []*NestedInteger, depth int)
    dfs = func(nl []*NestedInteger, depth int) {
        // base
        
        for _, n := range nl {
            if n.IsInteger() {
                weight := maxDepth-depth+1
                total += weight * n.GetInteger()
            } else {
                dfs(n.GetList(), depth+1)   
            }
        }
    }
    dfs(nestedList, 1)
    return total
}