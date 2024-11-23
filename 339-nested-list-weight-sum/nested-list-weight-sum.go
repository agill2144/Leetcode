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
// func depthSum(nestedList []*NestedInteger) int {
//     var dfs func(depth int, nl []*NestedInteger) int
//     dfs = func(depth int,  nl []*NestedInteger) int{
//         // base
//         if nl == nil {return 0}

//         // logic
//         total := 0
//         for i := 0; i < len(nl); i++{
//             if nl[i].IsInteger() {
//                 total += (nl[i].GetInteger() * depth)
//             } else {
//                 total += dfs(depth+1,nl[i].GetList())
//             }
//         }
//         return total
//     }
//     return dfs(1, nestedList)
// }


func depthSum(nestedList []*NestedInteger) int {
    total := 0
    type qNode struct {
        nl []*NestedInteger
        depth int
    }
    q := []*qNode{&qNode{nestedList, 1}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        nl := dq.nl
        depth := dq.depth
        for i := 0; i < len(nl); i++ {
            if nl[i].IsInteger() {
                total += nl[i].GetInteger() * depth
            } else {
                q = append(q, &qNode{nl[i].GetList(), depth+1})
            }
        }
    }
    return total
}