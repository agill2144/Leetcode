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
    out := new(strings.Builder)
    q := []*TreeNode{root}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq != nil {
            out.WriteString(fmt.Sprintf("%v",dq.Val))
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        } else {
            out.WriteString("#")
        }
        out.WriteString(",")
    }
    return out.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if len(data) == 0 {return nil}
    dataList := strings.Split(data, ",")
    ptr := 1
    rootVal, _ := strconv.Atoi(dataList[0])
    root := &TreeNode{Val: rootVal}
    q := []*TreeNode{root}
    for len(q) != 0 && ptr < len(dataList)  {
        dq := q[0]
        q = q[1:]
        leftChildVal := dataList[ptr]
        ptr++
        if leftChildVal == "#" {
            dq.Left = nil
        } else if leftChildVal != "#" {
            intVal, _ := strconv.Atoi(leftChildVal)
            dq.Left = &TreeNode{Val: intVal}
            q = append(q, dq.Left)
        }
        if ptr < len(dataList) {
            rightChildVal := dataList[ptr]
            ptr++
            if rightChildVal == "#" {
                dq.Right = nil
            } else if rightChildVal != "#" {
                intVal, _ := strconv.Atoi(rightChildVal)
                dq.Right = &TreeNode{Val: intVal}
                q = append(q, dq.Right)
            }
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


 4-7-3##-9-39-7-4#6#-6-6##065#9##-1-4###-2#######

 */