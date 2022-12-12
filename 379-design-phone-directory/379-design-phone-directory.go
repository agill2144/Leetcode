type PhoneDirectory struct {
    q *queue
    reserved map[int]struct{}
    max int
}


func Constructor(maxNumbers int) PhoneDirectory {
    q := newQueue(maxNumbers)
    return PhoneDirectory{
        q: q,
        max: maxNumbers,
        reserved: map[int]struct{}{},
    }
}


func (this *PhoneDirectory) Get() int {
    if len(this.reserved) == this.max {
        return -1
    }

    dq := this.q.dequeue()
    this.reserved[dq] = struct{}{}
    return dq
}

// Returns true if the number is available and false otherwise.
func (this *PhoneDirectory) Check(number int) bool {
    _, reserved := this.reserved[number]
    return !reserved
}


func (this *PhoneDirectory) Release(number int)  {
    if _, ok := this.reserved[number]; ok {
        delete(this.reserved, number)
        this.q.enqueue(number)
    }

}


/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */


/*** Queue Impl ***/
type listNode struct {
    val int
    next *listNode
}

type queue struct {
    head *listNode
    tail *listNode
}

func newQueue(maxNums int) *queue {
    head := &listNode{val: 0}
    tail := head
    // 0->1->2(h)
    for i := 1; i < maxNums; i++ {
        tail.next = &listNode{val: i}
        tail = tail.next
    }
    return &queue{
        head: head,
        tail: tail,
    }
}

func (q *queue) enqueue(x int) {
    newNode := &listNode{val: x}
    if q.head == nil {
        q.head = newNode
        q.tail = newNode
        return
    }
    q.tail.next = newNode
    q.tail = newNode
}


// 1(h)-2-3-4
func (q *queue) dequeue() int {
    if q.head == nil {
        return -1
    }
    // 1(h)-2-3-4
    out := q.head.val
    newHead := q.head.next
    q.head.next = nil
    q.head = newHead
    return out
}

func printLL(head *listNode) {
    curr := head
    msg := ""
    for curr != nil {
        msg += fmt.Sprintf("%v", curr.val)
        if curr.next != nil {
            msg += fmt.Sprintf("->")
        }
        curr = curr.next
    }
    fmt.Println(msg)
}