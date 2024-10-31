var outterArr = 1001
var innerArr = outterArr+1
type MyHashSet struct {
    items [][]bool
}


func Constructor() MyHashSet {
    return MyHashSet{items: make([][]bool, outterArr)}
}


func (this *MyHashSet) Add(key int)  {
    idx1 := hashIdx(key)
    idx2 := hashIdx2(key)
    if this.items[idx1] == nil {
        this.items[idx1] = make([]bool, innerArr)
    }
    this.items[idx1][idx2] = true
}

func hashIdx(key int) int {
    return key % outterArr
}

func hashIdx2(key int) int {
    return key % innerArr
}


func (this *MyHashSet) Remove(key int)  {
    idx1 := hashIdx(key)
    idx2 := hashIdx2(key)
    if this.items[idx1] == nil { return }
    this.items[idx1][idx2] = false
}


func (this *MyHashSet) Contains(key int) bool {
    idx1 := hashIdx(key)
    idx2 := hashIdx2(key)
    if this.items[idx1] == nil { return false }
    return this.items[idx1][idx2]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */