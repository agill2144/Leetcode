/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

/*
    Serialization is the process of converting a data structure or object into a sequence of bits(ie string)
    Deserialization is the process of converting sequence of bits(ie string) of data back into original data structure or object
    
    approach:
    - serialize: BFS
        - We need to follow a specific path of the input tree
        - And when you closely look, we are traversing level by level from left to right
        - Therefore we can use BFS with a queue to solve this problem
        - How?
        - We will obviously use a string builder for this
        - Then we will enqueue our root node
        - When processing the queue, we will check
        - Whether the current dq'd item is nil
        - Why? 
            - Because we will enqueue nil childs so that we can add null to our output string in the correct place
            - Wouldnt that end up adding extra nulls at the end , after the last level?
            - Yes, but that is still considered a successful serialization
        - if the dq'd item is nil
            - then simply write "null" to out string builder
        - otherwiese dq'd item is not nil
            - then simply write the value of the dq'd node to our string builder
            - BLINDLY ADD LEFT AND RIGHT CHILDS ( EVEN IF THEY ARE NILL )
            - Because we need "null" values in the middle of string, the extra nulls from the last level does not matter
        - This way we are able to convert out obj to string while keeping correct sequence of nulls
    - deserialize: BFS
        - First we need to convert our input string to something more parsable and iterable ( we can hack with pointers )
            - but no need to try and save space since we already will use a queue
            - so split the input string at ","
        - The idea is to construct the tree in a level order fashion
        - We know for sure the idx 0 value is our root of the tree, so...
        - First create the initial root node ( value will be idx 0 ofc )
        - Enqueue that root node into bfs queue
        - And have a pointer pointing to idx 1
        - For each level, we will dq a node
        - Then check if idx has a node to be added to current dq'd left child
        - (i.e value at idx is not "null")
        - if true ^ - then simply attach a left child to dq'd node and enqueue the left child
            - since this left child may have more childs to add
            - if not, the input string will tell us -- represented as "null" in input string
            - which we will ignore completely if the idx is pointing to "null"
        - move the idx forward
        - Apply the same for right child (make sure idx is not out of bounds )
        - move the idx forward
        - So our idx moves by twice when processing a node from queue
            - and it makes sense because a node has 2 childs
            - it either has it or does not ( null vs actual value in our input string )
        - All of this while our queue is not empty and while our idx still has items to process/add
*/


// Serializes a tree to a single string.
// time: o(n)
// space: o(n) 
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    str := new(strings.Builder)
    q := []*TreeNode{root}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == nil {
            str.WriteString("null")
        } else {
            str.WriteString(fmt.Sprintf("%v", dq.Val))
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        }
        str.WriteString(",")
    }
    
    return str.String()
}

// serialize using level order using dfs
// func (this *Codec) serialize(root *TreeNode) string {
//     if root == nil {
//         return ""
//     }
//     levels := [][]string{}
//     var dfs func(r *TreeNode, level int)
//     dfs = func(r *TreeNode, level int) {
//         // base
        
//         // logic
//         if len(levels) == level {
//             levels = append(levels, []string{})
//         }
//         levelStr := "null"
//         if r != nil {
//             levelStr = fmt.Sprintf("%v", r.Val)
//         }
//         levels[level] = append(levels[level], levelStr)
//         if r != nil {
//             dfs(r.Left, level+1)
//             dfs(r.Right, level+1)
//         }
//     }
//     dfs(root, 0)
//     strBuilder := new(strings.Builder)
//     for _, level := range levels {
//         strBuilder.WriteString(strings.Join(level, ","))
//         strBuilder.WriteString(",")
//     }
//     return strBuilder.String()
// }



// Deserializes your encoded data to tree.
// time: o(n)
// space: o(n) 
func (this *Codec) deserialize(data string) *TreeNode {    
    if data == "" {return nil}
    strData := strings.Split(data, ",")
    intVal,_ := strconv.Atoi(strData[0])
    root := &TreeNode{Val: intVal}
    idx := 1
    q := []*TreeNode{root}
    
    for len(q) != 0 && idx < len(strData) {
        dq := q[0]
        q = q[1:]
        if strData[idx] != "null" {
            intVal,_ := strconv.Atoi(strData[idx])
            dq.Left = &TreeNode{Val: intVal}
            q = append(q, dq.Left)
        }
        idx++
        if idx < len(strData) && strData[idx] != "null" {
            intVal,_ := strconv.Atoi(strData[idx])
            dq.Right = &TreeNode{Val: intVal}
            q = append(q, dq.Right)
        }
        idx++
    }
    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */