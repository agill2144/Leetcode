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

type PeekingIterator struct {
    iter *Iterator
    n int
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter: iter, n: iter.next()}
}

func (this *PeekingIterator) hasNext() bool {
    return this.n != -1
}

func (this *PeekingIterator) next() int {
    out := this.n
    if this.iter.hasNext() {
        this.n = this.iter.next()
    } else {
        this.n = -1
    }
    return out
}

func (this *PeekingIterator) peek() int {
    return this.n
}