const outterArraySize = 1000
const innerArraySize = 1001

/*
    TODO:
    1. Why is inner array size bigger by just 1 ?
    2. Why is our double hashing func using / operator instead of % operator ( like the 1st hashing func )
*/

func hashIdx(key int) int {  return key % outterArraySize }
func hashIdx2(key int) int { return key / innerArraySize  }


type MyHashSet struct {
    items [][]bool
}

func Constructor() MyHashSet {
    items := make([][]bool, 1000)
    return MyHashSet{
        items: items,
    }    
}


func (this *MyHashSet) Add(key int)  {
    idx := hashIdx(key)
    idx2 := hashIdx2(key)
    if this.items[idx] == nil {
        this.items[idx] = make([]bool, innerArraySize)
    }
    this.items[idx][idx2] = true
}


func (this *MyHashSet) Remove(key int)  {
    idx := hashIdx(key)
    idx2 := hashIdx2(key)
    if this.items[idx] == nil {
        return
    }
    this.items[idx][idx2] = false
}


func (this *MyHashSet) Contains(key int) bool {
    idx := hashIdx(key)
    idx2 := hashIdx2(key)
    if this.items[idx] == nil {
        return false
    }
    return this.items[idx][idx2]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */