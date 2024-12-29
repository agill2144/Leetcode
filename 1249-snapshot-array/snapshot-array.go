type SnapshotArray struct {
    snapId int
    history [][][]int // [ [], [] ]
}


func Constructor(length int) SnapshotArray {
    return SnapshotArray{
        history: make([][][]int, length),
    } 
}


func (this *SnapshotArray) Set(index int, val int)  {
    pair := []int{val, this.snapId}
    if this.history[index] == nil {
        this.history[index] = [][]int{}
    }
    this.history[index] = append(this.history[index], pair)
}


func (this *SnapshotArray) Snap() int {
    out := this.snapId
    this.snapId++
    return out
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    items := this.history[index]
    left := 0
    right := len(items)-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        if items[mid][1] <= snap_id {
            ans = items[mid][0]
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