/*
    len = 3
                    0                               1                           2
    [ [] , [[val, snapID], [val, snapID]], [[val, snapID], [val, snapID]]]
*/
type SnapshotArray struct {
    history [][][]int    
    id int
}


func Constructor(length int) SnapshotArray {
    data := make([][][]int, length)
    return SnapshotArray{
        history: data,
        id: 0,
    }
}


func (this *SnapshotArray) Set(index int, val int)  {
    records := this.history[index]    
    if records == nil {records = [][]int{}}
    records = append(records, []int{val, this.id})
    this.history[index] = records
    // fmt.Println("setting idx: ", index, " to val: ", val, " curr snapID: ", this.id)
    // fmt.Println(this.history)
    // fmt.Println("********************")
}


func (this *SnapshotArray) Snap() int {
    out := this.id
    this.id++
    return out    
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    records := this.history[index]
    left := 0
    right := len(records)-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        if records[mid][1] <= snap_id {
            ans = records[mid][0]
            left = mid+1
            // if records[mid][1] == snap_id{break}
        } else if snap_id > records[mid][1] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    // fmt.Println("get idx: ", index, " at snapID: ", snap_id, records)
    // fmt.Println("ans: ", ans)
    // fmt.Println("********************")
    return ans

}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */