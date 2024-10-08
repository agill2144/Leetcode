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
    if root == nil {return ""}
    out := &strings.Builder{}    
    q := []*TreeNode{root}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == nil {
            out.WriteString("#")
        } else {
            out.WriteString(fmt.Sprintf("%v", dq.Val))
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        }
        // use "," as delim cuz there are negative values ( instead of - as delim )
        out.WriteString(",")
    }

    return out.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    dataSplit := strings.Split(data,",")
    dataSplit = dataSplit[:len(dataSplit)-1]
    if len(dataSplit) == 0 {return nil}
    rootVal, _ := strconv.Atoi(dataSplit[0])
    root := &TreeNode{Val: rootVal}
    ptr := 1
    q := []*TreeNode{root}
    for len(q) !=0 && ptr < len(dataSplit) {
        dq := q[0]
        q = q[1:]

        // left child
        ptrVal := dataSplit[ptr]
        ptr++
        if ptrVal == "#" {
            dq.Left = nil
        } else {
            intVal, _ := strconv.Atoi(ptrVal)
            dq.Left = &TreeNode{Val: intVal}
            q = append(q, dq.Left)
        }

        if ptr == len(dataSplit) {break}

        // right child
        ptrVal = dataSplit[ptr]
        ptr++
        if ptrVal == "#" {
            dq.Right = nil
        } else {
            intVal, _ := strconv.Atoi(ptrVal)
            dq.Right = &TreeNode{Val: intVal}
            q = append(q, dq.Right)
        }
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