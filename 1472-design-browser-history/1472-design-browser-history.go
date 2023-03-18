type BrowserHistory struct {
    dll *dll
    curr *listNode
}


func Constructor(homepage string) BrowserHistory {
    dll := new(dll)
    curr := dll.addToEnd(homepage)
    return BrowserHistory{dll: dll, curr: curr}
}


func (this *BrowserHistory) Visit(url string)  {
    this.dll.removeEverythingAfterNode(this.curr)
    this.curr = this.dll.addToEnd(url)
}


func (this *BrowserHistory) Back(steps int) string {
    for steps != 0 && this.curr != this.dll.head {
        this.curr = this.curr.prev
        steps--
    } 
    return this.curr.val
}


func (this *BrowserHistory) Forward(steps int) string {
    for steps != 0 && this.curr != this.dll.tail {
        this.curr = this.curr.next
        steps--
    } 
    return this.curr.val
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */


type listNode struct {
    val string
    next *listNode
    prev *listNode
}

type dll struct {
    head *listNode
    tail *listNode
}


func (d *dll) addToEnd(val string) *listNode {
    newNode := &listNode{val: val}
    if d.head == nil {
        d.head = newNode
        d.tail = newNode
        return newNode
    }
    newNode.prev = d.tail
    d.tail.next = newNode
    d.tail = newNode
    return newNode
}


func (d *dll) removeEverythingAfterNode(n *listNode) {
    if n == d.tail {return}
    if n == d.head {
        if d.head.next != nil {d.head.next.prev = nil}
        d.head.next = nil    
        d.tail = d.head
        return
    }
    n.next.prev = nil
    n.next = nil
    d.tail = n
}



// array implementation
// arrays = cpu cache, so somewhat performant compared to LL
// type BrowserHistory struct {
//     items []string
//     curr int
// }


// func Constructor(homepage string) BrowserHistory {
//     return BrowserHistory{
//         items: []string{homepage},
//         curr: 0,
//     }
// }


// func (this *BrowserHistory) Visit(url string)  {
//     if this.curr != len(this.items) -1 {
//         this.items = this.items[:this.curr+1]
//     }
//     this.items = append(this.items, url)
//     this.curr++
// }


// func (this *BrowserHistory) Back(steps int) string {
//     for steps != 0 && this.curr != 0 {
//         this.curr--
//         steps--
//     }
//     return this.items[this.curr]
// }


// func (this *BrowserHistory) Forward(steps int) string {
//     for this.curr != len(this.items)-1 && steps != 0 {
//         this.curr++
//         steps--
//     }
//     return this.items[this.curr]
// }


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
