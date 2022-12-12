
/*
    approach 1: inorder traversal with recursion
    
    - We will save local state at each node
        - instead of saving a running sum at each node, we will subtract target from root.val
        - So target will be maintained at each node ( updated target )
        - We will also keep track of a list thats maintaining running path ( using go slices )*
            - slices are pass by value
            - slices have len and cap as properties
            - cap = total size avail
            - len = current occupied size
            - slices also have a ptr to 1 underlying array and they use len and cap to make growing and shrinking decisions
            - this is important because since we wont allocate a specific size slice we want the slice to grow and shrink as needed
            - the side effect of this is using in recursion, is that at some node len will be 3 while cap will be 4 ( meaning the underlying array still has size )
            - remember the underlying arr is a ref.
            - slice is pass by value ( copy of top level meta info like len, cap, but arr ref is STILL THE SAME )
            - but when the slice says that it still has capacity , we go down a branch and come back and to the same node that said the underlying arr still has capacity
            - but! it does not. The local state is out of date because when we down a branch , it added an element and made len == cap
            - but when the recursion popped the top of the stack, in that local state for the same slice is something else ( because slices are pass by value )
            - so we have 2 different states while the underlying array may be full..
            - The only workaround I have found so far is - when saving the result, ENSURE its being copied into a new slice and then save the new slice or else 
                - the underlying array ( ref ) gets changed and then your result will change too....
            - in python, its all pass by ref, so we can easily backtrack and remove 1 element and all nodes have the same truth vs in Golang the truth once we go back to a prev node is a lie! so careful! ( I learned this after spending an entire day of troubleshooting .......)
        - once we run into a target == 0 and current node is a leaf node
        - Then we will add the running path to a resulting list
            - which means we will also maintain a reference to a slice of result array
        
        - time: o(n^2) - we visit every single node and when we do find an answer, we have to copy 1 slice to another to avoid the side effects above
        - space:
            - o(h) for recursion stack
            - and when we do find an answer, we allocate another o(n) slice to copy current paths slice into
            - o(hn)
            Not sure if this is correct ^
        
        
        approach 2: inorder but iterative
        - When we pop the top of the stack,
        - We must reset the "local recursion" state to match the top of the stack
        - this also had the same slice problem ( so had to copy )
        - time: o(n^2)
        - space: o(h for stack) * if at the current node we find an answer we allocate o(n) to save paths - o(hn)

*/



func pathSum(root *TreeNode, targetSum int) [][]int {
    var result [][]int
    inorderDfs(root, targetSum, nil, &result)
    return result
}

func inorderDfs(root *TreeNode, targetSum int, paths []int, result *[][]int) {
    
    
    // base ( since this is void func, we return and do nothing )
    if root == nil {
        return
    }
    
    // logic
    targetSum = targetSum - root.Val
    paths = append(paths, root.Val)
    inorderDfs(root.Left, targetSum, paths, result)
    
    if targetSum == 0 && root.Left == nil && root.Right == nil {
        newList := make([]int, len(paths))
        copy(newList, paths)
        *result = append(*result, newList)
    }
    inorderDfs(root.Right, targetSum, paths, result)
    paths = paths[:len(paths)-1]
}



// approach 2  : inorder but iterative

// type obj struct {
//     node *TreeNode
//     paths []int
//     sum int
// }

// func pathSum(root *TreeNode, targetSum int) [][]int{
//     if root == nil {return nil}
//     var result [][]int
//     stack := []*obj{}
//     paths := []int{}
//     runningSum := 0

//     for root != nil || len(stack) != 0 {
        
//         for root != nil {
//             runningSum += root.Val
//             paths = append(paths, root.Val)
//             stack = append(stack, &obj{node: root, paths: paths, sum: runningSum})
//             root = root.Left
//         }
//         o := stack[len(stack)-1]
//         stack = stack[:len(stack)-1]
//         p := o.paths
//         root = o.node
//         s := o.sum
//         // fmt.Println(root.Val, p, s)
//         if targetSum == s && root.Left == nil && root.Right == nil {
//             newList := make([]int, len(p))
//             copy(newList, p)
//             result = append(result, newList)
//         }
//         root = root.Right
//         // reset the same things recursion implicitly resets at the local state but explicitly here
//         runningSum = s
//         paths = p
//     }
//     return result
// }