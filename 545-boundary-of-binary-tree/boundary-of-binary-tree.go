/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
approach: naive, 3-4 passes
- collect left view
- collect leaves
- collect right
- reverse right
- return concat ( left + leaves + right )
tc = o(n)
sc = o(n) = recursion space

approach: top-down / preorder recursion, 2 passes
- marking nodes with an ID
- root = 0
- left boundary node = 1
- right boundary node = 2
- middle node = 3
- leaves are easy to detect, they have no left or right child, therefore it's ID does not matter!
- assume we have a node in hand, and we know its ID, and we are about to go left, what will be this left child ID ?
   - root's left will always be left boundary node ( if id == 0 { return 1 } )
   - if current node is a left boundary node then, its left will always be left boundary node ( if id == 1 {return 1}
   - if current node is a right boundary node then;
     - if this node does not have a right child, then the left node will be right right boundary
   - all other cases, this node's left will be a middle node
- assume we have a node in hand, and we know its ID, and we are about to go right, what will be this right child ID ?
   - root's right will always be right boundary node ( if id == 0 { return 2 } )
   - if current node is a right boundary node then, its right will always be right boundary node ( if id == 2 {return 2}
   - if current node is a left boundary node then;
     - if this node does not have a left child, then the right node will be a left boundary node
   - all other cases, this node's right will be a middle node

- in pre-order fashion, process root first!
- i.e process curr ID, if its a leaf node then saves to leaves arr, if 0 || 1 then save to left view arr, if 2 then save to right view array/stack
- then recurse left , with left child's ID ( parent will identify this value and pass down the value to the recursion )
- reucrse right, with right child's ID ( parent will identify this value and pass down the value to the recursion )
*/

/*
    0 = root
    1 = left
    2 = right
    3 = middle
*/

func boundaryOfBinaryTree(root *TreeNode) []int {
    left := []int{}
    leaves := []int{}
    right := []int{}
    var dfs func(r *TreeNode, id int)
    dfs = func(r *TreeNode, id int) {
        // base
        if r == nil {return}

        // logic
        if r.Left == nil && r.Right == nil {
            leaves = append(leaves, r.Val)
        } else if id == 0 || id == 1 {
            left = append(left, r.Val)
        } else if id == 2 {
            right = append(right, r.Val)
        }

        dfs(r.Left, leftID(r, id))
        dfs(r.Right, rightID(r, id))
    }
    dfs(root, 0)
    left = append(left, leaves...)
    for i := len(right)-1; i >= 0; i-- {
        left = append(left, right[i])
    }
    return left
}

func leftID(r *TreeNode, currID int) int {
    if currID == 0 || currID == 1 {
        return 1
    } else if currID == 2 {
        if r.Right == nil {
            return 2
        }
    }
    return 3
}
func rightID(r *TreeNode, currID int) int {
    if currID == 0 || currID == 2 {
        return 2
    } else if currID == 1 {
        if r.Left == nil {
            return 1
        }
    }
    return 3
}