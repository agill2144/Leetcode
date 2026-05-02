var (
    outter int = 1000
    inner int = 1001
)
type MyHashSet struct {
    data [][]bool
}


func Constructor() MyHashSet {
    return MyHashSet{
        data: make([][]bool, outter),
    }
}

func (this *MyHashSet) hash1(key int) int {
    return key % outter
}

func (this *MyHashSet) hash2(key int) int {
    return key / inner
}


func (this *MyHashSet) Add(key int)  {
    idx1 := this.hash1(key)
    idx2 := this.hash2(key)
    if this.data[idx1] == nil {
        this.data[idx1] = make([]bool, inner)
    }
    this.data[idx1][idx2] = true
}


func (this *MyHashSet) Remove(key int)  {
    idx1 := this.hash1(key)
    idx2 := this.hash2(key)
    if this.data[idx1] == nil {return}
    this.data[idx1][idx2] = false
}


func (this *MyHashSet) Contains(key int) bool {
    idx1 := this.hash1(key)
    idx2 := this.hash2(key)
    if this.data[idx1] == nil {return false}
    return this.data[idx1][idx2]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */