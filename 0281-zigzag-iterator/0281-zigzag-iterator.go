type ZigzagIterator struct {
    q [][]int
    v1 []int
    v2 []int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
    q := [][]int{}
    if len(v1) != 0 {
        q = append(q, []int{1,0})
    }
    if len(v2) != 0 {
        q = append(q, []int{2,0})
    }
    return &ZigzagIterator{
        q: q, // [ [listID, ptr] ]
        v1: v1,
        v2: v2,
    }
}

func (this *ZigzagIterator) next() int {
    dq := this.q[0]
    this.q = this.q[1:]
    listId := dq[0]
    ptr := dq[1]
    var out int
    if listId == 1  {
        out = this.v1[ptr]
        ptr++
        if ptr < len(this.v1) {
            this.q = append(this.q, []int{listId, ptr})
        }
    } else  {
        out = this.v2[ptr]
        ptr++
        if ptr < len(this.v2) {
            this.q = append(this.q, []int{listId, ptr})
        }
    }
    return out
}

func (this *ZigzagIterator) hasNext() bool {
    return len(this.q) != 0
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */