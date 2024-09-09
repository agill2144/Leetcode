type SnapshotArray struct {
    // {idx : {snapID: val}}
    items map[int]map[int]int
    snapID int
}


func Constructor(length int) SnapshotArray {
    v := map[int]map[int]int{}    
    for i := 0; i < length; i++ {
        v[i] = map[int]int{} // idx: {snapID: val}
    }
    return SnapshotArray{
        items: v,
        snapID: 0,
    }
}


func (this *SnapshotArray) Set(idx int, val int)  {
    this.items[idx][this.snapID] = val
}


func (this *SnapshotArray) Snap() int {
    this.snapID++
    return this.snapID-1    
}


func (this *SnapshotArray) Get(idx int, snap_id int) int {
    if _, ok := this.items[idx][snap_id]; ok {return this.items[idx][snap_id]}
    for snap_id > 0 {
        snap_id--
        _, ok := this.items[idx][snap_id]
        if ok {break}
    }
    if snap_id < 0 {return 0}
    return this.items[idx][snap_id]
}



/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */