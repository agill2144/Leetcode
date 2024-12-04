type SparseVector struct {
    items [][]int // [[val, idx], [val2, idx2]]
}

func Constructor(nums []int) SparseVector {
    items := [][]int{}
    for i := 0; i < len(nums); i++{
        if nums[i] == 0{continue}
        items = append(items, []int{nums[i], i})
    }
    return SparseVector{items}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    sum := 0
    v1, v2 := 0, 0
    for v1 < len(this.items) && v2 < len(vec.items) {
        v1Idx, v2Idx := this.items[v1][1], vec.items[v2][1]
        if v1Idx == v2Idx {
            sum += (this.items[v1][0] * vec.items[v2][0])
            v1++; v2++
        } else if v1Idx < v2Idx {
            v1++
        } else {
            v2++
        }
    }
    return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */