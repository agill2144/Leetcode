type SparseVector struct {
    // [ [idx, val], [idx, val] ] - only contains items with non-zero values
    items [][]int 
}

func Constructor(nums []int) SparseVector {
    items := [][]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {continue}
        items = append(items, []int{i, nums[i]})
    }
    return SparseVector{items}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    p1 := 0
    p2 := 0
    res := 0
    for p1 < len(this.items) && p2 < len(vec.items) {
        idx1 := this.items[p1][0]
        idx2 := vec.items[p2][0]
        if idx1 == idx2 {
            res += (this.items[p1][1] * vec.items[p2][1])
            p1++; p2++
        } else if idx1 < idx2 {
            p1++
        } else {
            p2++
        }
    }
    return res
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */