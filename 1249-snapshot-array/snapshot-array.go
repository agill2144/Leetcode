
type SnapshotArray struct {
    id int
    histories [][][]int
}


func Constructor(length int) SnapshotArray {
    return SnapshotArray{
        histories: make([][][]int, length),
        id: 0,
    }
}


func (this *SnapshotArray) Set(index int, val int)  {
    vals := this.histories[index]
    if vals == nil {vals = [][]int{}}
    prevID := -1
    if len(vals) != 0 {prevID = vals[len(vals)-1][1]}
    if prevID == this.id {
        vals[len(vals)-1][0] = val
    } else {
        vals = append(vals, []int{val, this.id})
    }
    this.histories[index] = vals
}


func (this *SnapshotArray) Snap() int {
    out := this.id
    this.id++
    return out
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    vals := this.histories[index]        
    left := 0
    right := len(vals)-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        if vals[mid][1] <= snap_id {
            ans = vals[mid][0]
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */