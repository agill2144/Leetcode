/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

// we care that the top is the next number
// whenever we call next, save the current top, call iter.Next() to save next top element

type PeekingIterator struct {
    top int
    iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
    out := &PeekingIterator{
        top: -1,
        iter: iter,
    }
    out.ensureTopIsSet()
    return out
}

func (this *PeekingIterator) ensureTopIsSet() {
    if this.iter.hasNext() {
        this.top = this.iter.next()
    } else {
        // no more elements left 
        this.top = -1
    }
}

func (this *PeekingIterator) hasNext() bool {
    return this.top != -1
}

func (this *PeekingIterator) next() int {
    out := this.top
    this.ensureTopIsSet()
    return out
}

func (this *PeekingIterator) peek() int {
    return this.top    
}