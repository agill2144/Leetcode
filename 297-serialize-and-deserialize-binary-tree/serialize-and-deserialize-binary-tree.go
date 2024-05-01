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

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    sb := new(strings.Builder)
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        sb.WriteString(fmt.Sprintf("%v,", r.Val))
        if r.Left == nil {sb.WriteString("#,")}
        dfs(r.Left)
        if r.Right == nil {sb.WriteString("#,")}
        dfs(r.Right)
    }
    dfs(root)
    return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    if len(data) == 0 {return nil}
    split := strings.Split(data, ",")  
    split = split[:len(split)-1]
    ptr := 0
    var dfs func() *TreeNode
    dfs = func () *TreeNode {
        // base
        if ptr == len(split) {return nil}

        // logic
        if split[ptr] == "#" {ptr++; return nil}
        intVal, _ := strconv.Atoi(split[ptr])
        root := &TreeNode{Val: intVal}
        ptr++
        root.Left = dfs()
        root.Right = dfs()
        return root
    }
    return dfs()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */