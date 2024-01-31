func largestPalindromic(num string) string {
    freq := make([]int, 11) // idx -> num mapping
    for i := 0; i < len(num); i++ {
        char := num[i]
        freq[int(char-'0')]++
    }
    out := new(strings.Builder)
    var middle string
    back := []string{}
    for i := len(freq)-1; i >= 0; i-- {
        if freq[i] <= 1 {
            continue
        }
        count := freq[i] / 2
        freq[i] = freq[i] % 2
        for k := 0; k < count; k++ {
            if out.String() == "" && i == 0 {continue}
            out.WriteString(fmt.Sprintf("%v", i))
            back = append(back, fmt.Sprintf("%v", i))
        }
    }
    for i := len(freq)-1; i >= 0; i-- {
        if freq[i] >= 1 {
            middle = fmt.Sprintf("%v", i)
            break
        }
    }
    out.WriteString(middle)
    for i := len(back)-1; i >= 0; i-- {
        out.WriteString(back[i])
    }
    if out.String() == "" {out.WriteString("0")}
    return out.String()
    
}


type listNode struct {
    val int
    next *listNode
    prev *listNode
}
type linkedList struct {
    head *listNode
    tail *listNode
}
func (l *linkedList) addToHead(val int) {
    newHead := &listNode{val: val}
    if l.head == nil {
        l.head = newHead
        l.tail = newHead
        return
    }
    
    newHead.next = l.head
    l.head.prev = newHead
    l.head = newHead
}
func (l *linkedList) addToTail(val int) {
    newTail := &listNode{val: val}
    if l.head == nil {
        l.addToHead(val)
        return
    }
    l.tail.next = newTail
    newTail.prev = l.tail
    l.tail = newTail
}
func (l *linkedList) removeHead() int {
    out := l.head.val
    if l.head.next == nil {
        l.head = nil
        l.tail = nil
        return out
    }
    newHead := l.head.next
    l.head.next = nil
    newHead.prev = nil
    l.head = newHead
    return out
}
func (l *linkedList) removeTail() int {
    out := l.tail.val
    if l.head == l.tail {
        l.head = nil
        l.tail = nil
        return out
    }
    newTail := l.tail.prev
    newTail.next = nil
    l.tail.prev = nil
    l.tail = newTail
    return out
}
