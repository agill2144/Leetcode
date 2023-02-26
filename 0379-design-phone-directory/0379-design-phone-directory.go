type PhoneDirectory struct {
    max int
    sll *sll
    taken map[int]struct{}
}


func Constructor(maxNumbers int) PhoneDirectory {
    sll := &sll{}
    for i := 0; i < maxNumbers; i++ {
        sll.Append(i)
    }
    return PhoneDirectory{
        max: maxNumbers,
        sll: sll,
        taken: map[int]struct{}{},
    }
}

// time: o(1)
// space: o(1)
func (this *PhoneDirectory) Get() int {
    if this.sll.Size() == 0 { return -1 }
    out := this.sll.PopFromFront()
    this.taken[out.val] = struct{}{}
    return out.val
}

// time: o(1)
// space: o(1)
func (this *PhoneDirectory) Check(number int) bool {
    _, ok := this.taken[number]
    return !ok
}

// time: o(1)
// space: o(1)
func (this *PhoneDirectory) Release(number int)  {
    _, ok := this.taken[number]
    if ok {
        this.sll.Append(number)
        delete(this.taken, number)
    }
}


type listNode struct {
    val int
    next *listNode
}

type sll struct {
    head *listNode
    tail *listNode
    size int
}

// time: o(1)
// space: o(1)
func (s *sll) Size() int { return s.size }

// time: o(1)
// space: o(1)
func (s *sll) Append(x int) {
    newTail := &listNode{val: x}
    if s.head == nil {
        s.head = newTail
        s.tail = newTail
        s.size = 1
        return
    }
    s.tail.next = newTail
    s.tail = newTail
    s.size++
}

// time: o(1)
// space: o(1)
func (s *sll) PopFromFront() *listNode {
    out := s.head
    if s.head == nil || s.head.next == nil {
        s.size = 0
        s.head = nil
        s.tail = nil
        return out
    }
    newHead := s.head.next
    s.head.next = nil
    s.head = newHead
    s.size--
    return out
}

/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */