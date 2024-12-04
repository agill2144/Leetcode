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
    if len(vec.items) < len(this.items) {return vec.dotProduct(*this)}
    sum := 0
    for i := 0; i < len(this.items); i++ {
        val, idx := this.items[i][0], this.items[i][1]
        idx2 := search(vec.items, idx)
        if idx2 != -1 {
            sum += (val * vec.items[idx2][0])
        }
    }
    return sum
}

func search(items[][]int, targetIdx int) int {
    left := 0
    right := len(items)-1
    for left <= right {
        mid := left + (right-left)/2
        if items[mid][1] == targetIdx {return mid}
        if targetIdx > items[mid][1] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */