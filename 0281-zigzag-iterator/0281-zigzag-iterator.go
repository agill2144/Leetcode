type ZigzagIterator struct {
    v1, v2 []int
    p1, p2 int
    p1Moved bool
}

func Constructor(v1, v2 []int) *ZigzagIterator {
    return &ZigzagIterator{
        v1: v1, v2 : v2,
        p1: 0, p2 : 0,
        p1Moved: false,
    }
}

func (this *ZigzagIterator) moveP1() int {
    out := this.v1[this.p1]
    this.p1++
    this.p1Moved = true
    return out    
}
func (this *ZigzagIterator) moveP2() int {
    out := this.v2[this.p2]
    this.p2++
    this.p1Moved = false
    return out    
}

func (this *ZigzagIterator) next() int {
    if !this.p1Moved {
        if this.p1 < len(this.v1) {
            return this.moveP1()
        }
        return this.moveP2()
    } else {
        if this.p2 < len(this.v2) {
            return this.moveP2()
        }
        return this.moveP1()
    }
}

func (this *ZigzagIterator) hasNext() bool {
    return this.p1 < len(this.v1) || this.p2 < len(this.v2)
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */