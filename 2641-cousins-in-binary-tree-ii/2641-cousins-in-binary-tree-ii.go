/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: single pass
    - we need total level sum
    - we need each sibling sum
    - we know the level sum at level 0 is the root value itself
    - we do not know the next level by first taking a pass just to get the level sum of each level.
    - turns out we can do this while running a BFS
    - have each parent node do 2 things ( at each level )
        1. Look down at its child and get their child sum
            - this is the sibling sum for each of its child
            - and then the parent enqueue these childs with their siblingSum ; a pair < node, siblingSum >
            - in next level, when these childs are dq'd we would have the exact sibling sum to remove from total level sum
            - now we need the total level sum
        2. At each level, when we are computing the sibling sum by looking down, 
            - we can keep adding these sum to a temp var
            - once a level is finished, we promote the this temp value to a $levelSum var
            - so that when next level starts, and we dq a node, the dq'd node will have the exact sibling sum 
            - and the $level sum var will have this levels total sum
            - now we can replace the dq'd nodes value with its cousins sum ; $levelSum - $siblingSum
            - And when this node is updated, this node is also a parent of some childs, repeat #1 and #2

    time = o(n)
*/

func replaceValueInTree(root *TreeNode) *TreeNode {
    if root == nil {return nil}

    type queueNode struct{
        node *TreeNode
        sum int // siblingSum
    }
    levelSum := root.Val
    q := []*queueNode{&queueNode{root,root.Val}}
    for len(q) != 0 {
        nextLevelSum := 0
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq.node
            node.Val = levelSum-dq.sum
            
            nextSiblingSum := 0
            // why cant we also enqueue left node and right node here?
            // we cant, because with each node, we need to set its total sibling sum
            // which is left+right, thats why we figure the sibling sum first, then only we add it
            if node.Left != nil {nextSiblingSum += node.Left.Val}
            if node.Right != nil {nextSiblingSum += node.Right.Val}
            
            // add sibling sum for next level sum
            nextLevelSum += nextSiblingSum
            
            // enqueue childs with their exact sibling sum
            if node.Left != nil {
                q = append(q, &queueNode{node.Left,nextSiblingSum})
            }
            if node.Right != nil {
                q = append(q, &queueNode{node.Right,nextSiblingSum})
            }
            qSize--
        }
        // level is finished
        // promote nextLevelSum to levelSum 
        levelSum = nextLevelSum
    }
    return root
}


/*
    We have to replace a node value with sum of its cousins value
    Cousins are all nodes on the same level with different parents
    
    approach: two pass dfs
    - 1st pass: get level sum for each level and store it by level ( hashmap )
    - 2nd pass;
        - have each node look down and replace values of its childs
        - Each node has 2 childs, i.e this node is a parent of left and right child
        - so this node can look down and sum its childs sum ( siblings )
        - then this same node can get the next level sum ( currLevel + 1 ) from hashmap
        - and remove its childs sum from the levelSum
        - this becomes the new value for both of its childs
    
    time = o(n)
    space = o(n)
*/
// func replaceValueInTree(root *TreeNode) *TreeNode {
//     if root == nil {return root}
//     levelSum := map[int]int{}
//     var dfs1 func(r *TreeNode, level int)
//     dfs1 = func(r *TreeNode, level int) {
//         // base
//         if r == nil {return}
        
//         // logic
//         levelSum[level] += r.Val
//         dfs1(r.Left, level+1)
//         dfs1(r.Right, level+1)
//     }
//     // populate level sum
//     dfs1(root, 0)
    
//     var dfs2 func(r *TreeNode, level int)
//     dfs2 = func(r *TreeNode, level int) {
//         // base
//         if r == nil {return}
        
//         // logic
//         nextLevelSum, ok := levelSum[level+1]
//         if ok {
//             childSum := 0
//             if r.Left != nil {childSum += r.Left.Val}
//             if r.Right != nil {childSum += r.Right.Val}
//             cousinSum := nextLevelSum - childSum
//             if r.Left != nil {r.Left.Val = cousinSum}
//             if r.Right != nil {r.Right.Val = cousinSum}
//         }
//         dfs2(r.Left, level+1)
//         dfs2(r.Right, level+1)        
//     }
//     dfs2(root, 0)
//     root.Val = 0
//     return root
// }