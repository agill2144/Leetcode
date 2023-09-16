/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/*
  a root is a node;
  - does not have a parent
  - which also means , its never a child of some other node
  
  So we need to find a node that is not a child
  
  Approach: using a hashset to collect all childs
  - we have access to all nodes
  - for each node, loop over all the childs and save them in a childSet
    - i.e these are nodes which are childs of some parent
    - therefore these nodes are def not root
    - because a root is not a child of anyone
  - then, go thru the input nodes one more time
    - and check if this node exists as a child in childSet
    - if it does not, we have found the root node
    
time = o(2n)
space = o(n)
*/

func findRoot(tree []*Node) *Node {
    
    childSet := map[*Node]struct{}{}
    for _, node := range tree {
        for _, child := range node.Children {
            childSet[child] = struct{}{}
        }
    }

    for _, node := range tree {
        if _, ok := childSet[node]; !ok {
            return node
        }
    }
    return nil
}
