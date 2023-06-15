/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    if root == nil {return 0}
    nums := []int64{}
    q := []*TreeNode{root}
    for len(q) != 0 {
        qSize := len(q)
        var levelSum int64
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            levelSum += int64(dq.Val)
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        nums = append(nums, levelSum)
    }
    
    targetIdx := len(nums)-k
    if targetIdx < 0 || targetIdx == len(nums) {return -1}
    l := 0
    r := len(nums)-1
    for l <= r {
        
        nextSmaller := l
        pivotIdx := r
        for i := l; i < pivotIdx; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[i], nums[nextSmaller] = nums[nextSmaller], nums[i]
                nextSmaller++
            }
        }
        nums[nextSmaller], nums[pivotIdx] = nums[pivotIdx], nums[nextSmaller]
        if targetIdx == nextSmaller {
            return nums[nextSmaller]
        }
        
        if targetIdx > nextSmaller {
            l = nextSmaller+1
        } else {
            r = nextSmaller-1
        }
        
    }
    return nums[targetIdx]
}
