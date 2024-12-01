type SnapshotArray struct {
    items [][][]int
    id int
}


func Constructor(length int) SnapshotArray {
    return SnapshotArray{
        items: make([][][]int, length),
        id: 0,
    }
}


func (this *SnapshotArray) Set(index int, val int)  {
    items := this.items[index]
    if items == nil { items = [][]int{}}
    if len(items) > 0 && items[len(items)-1][1] == this.id {
        items[len(items)-1][0] = val
    } else {
        items = append(items, []int{val, this.id})

    }
    this.items[index] = items
}


func (this *SnapshotArray) Snap() int {
    this.id++
    return this.id-1
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    // right most on left side of snap_id
    items := this.items[index]
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

["SnapshotArray","set","snap","snap","snap","get","snap","snap","get"]
[[1],[0,15],[],[],[],[0,2],[],[],[0,0]]

SnapshotArray(1) ; id = 0
set(0,15)
snap() id = 1
snap() id = 2
snap() id = 3
get(0,2)
snap() id = 4
snap() id = 5
get(0,0) 
 */