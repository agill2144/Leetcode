// have each idx keep track of values it has had over each snapID
// snapID will incrementally increase
// therefore each idx will have a list of values [snapID, val]
// that tells what that value was at the snapID
// snapIDs in this nested list will be sorted, because they increase incrementally
// now we can use binary search to find the snapID in logN time
type SnapshotArray struct {
    history [][][]int
    id int
}


func Constructor(length int) SnapshotArray {
    return SnapshotArray{
        history: make([][][]int, length),
        id: 0,
    }
}


func (this *SnapshotArray) Set(index int, val int)  {
    if this.history[index] == nil {
        this.history[index] = [][]int{}
    }
    this.history[index] = append(this.history[index], []int{this.id, val})
}


func (this *SnapshotArray) Snap() int {
    out := this.id
    this.id++
    return out
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    history := this.history[index]
    left := 0
    right := len(history)-1
    val := 0
    for left <= right {
        mid := left + (right-left)/2
        if history[mid][0] <= snap_id {
            val = history[mid][1]
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return val
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */