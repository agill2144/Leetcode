func furthestBuilding(heights []int, bricks int, ladders int) int {
    mx := &maxHeap{items: []int{}}
    for i := 0; i < len(heights)-1; i++ {
        curr := heights[i]
        next := heights[i+1]
        if next <= curr {continue}

        
        diff := next-curr
        // always use bricks
        bricks -= diff
        // push the number of bricks used
        heap.Push(mx, diff)

        // if we used more bricks than what we had
        if bricks < 0 {
            // see if we can re-write history
            // i.e replace previous brick usage ( the highest bricks usage that sitting in max heap )
            // with a ladder

            // but if we do not have any ladders , return curr building position
            if ladders == 0 {
                return i
            }

            // we have ladders, re-write max-bricks-usage climb with a ladder instead
            ladders--
            // collect the bricks from last usage since it got replaced with ladder
            bricks += heap.Pop(mx).(int)
        }
    }
    return len(heights)-1
}

type maxHeap struct {items []int}
func (m *maxHeap) Less(i,j int) bool {return m.items[i] > m.items[j]}
func (m *maxHeap) Swap(i,j int) {m.items[i],m.items[j] = m.items[j],m.items[i]}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

// tried greedy, didnt work
// func furthestBuilding(heights []int, bricks int, ladders int) int {
//     b,l := bricks, ladders
//     i := 0
//     for i < len(heights)-1 {
//         curr := heights[i]
//         next := heights[i+1]
//         if next <= curr {i++;continue}
//         heightDiff := next-curr
//         if heightDiff <= b {
//             b -= heightDiff
//         } else if l >= 1 {
//             l--
//         } else {
//             break
//         }
//         i++
//     }
//     j := 0
//     for j < len(heights)-1 {
//         curr := heights[j]
//         next := heights[j+1]
//         // fmt.Println(i, curr, next, next-curr)
//         if next <= curr {j++;continue}
//         heightDiff := next-curr
//         if ladders >= 1 {
//             ladders--
//         } else if heightDiff <= bricks {
//             bricks -= heightDiff
//         } else {
//             break
//         }
//         j++
//     }

//     return max(i, j)
// }

/*
    [1,5,1,2,3,4,10000]
             i
*/