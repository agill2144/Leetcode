
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) == 0 || len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	cardCount := make(map[int]int)
	for _, card := range hand {
		cardCount[card]++
	}

	for _, card := range hand {
		if cardCount[card] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if cardCount[card+i] == 0 {
				return false
			}
			cardCount[card+i]--
		}
	}

	return true

}
// func isNStraightHand(hand []int, groupSize int) bool {
//     n := len(hand)
//     if n % groupSize != 0 {return false}
//     mh := &minHeap{items: []int{}}
//     freq := map[int]int{}
//     for i := 0; i < len(hand); i++ {
//         if _, ok := freq[hand[i]]; !ok {
//             heap.Push(mh, hand[i])
//         }
//         freq[hand[i]]++
//     }

//     grpSize := 0 
//     lastGrpNum := -1
//     pb := []int{} 
//     for mh.Len() != 0 {
//         popped := heap.Pop(mh).(int)
//         if grpSize == 0 {
//             lastGrpNum = popped
//             grpSize = 1
//         } else {
//             if lastGrpNum+1 != popped {return false}
//             lastGrpNum = popped
//             grpSize++
//         }

//         freq[popped]--
//         if freq[popped] > 0 {
//             pb = append(pb, popped)
//         }

//         if grpSize == groupSize {
//             grpSize = 0
//             for k := 0; k < len(pb); k++ {heap.Push(mh, pb[k])}
//             pb = []int{}
//         }
//     }

//     return grpSize == 0 && mh.Len() == 0
// }

// type minHeap struct {
// 	items []int
// }

// func (m *minHeap) Less(i, j int) bool {
// 	return m.items[i] < m.items[j]
// }
// func (m *minHeap) Swap(i, j int) {
// 	m.items[i], m.items[j] = m.items[j], m.items[i]
// }
// func (m *minHeap) Len() int {
// 	return len(m.items)
// }
// func (m *minHeap) Push(x interface{}) {
// 	m.items = append(m.items, x.(int))
// }
// func (m *minHeap) Pop() interface{} {
// 	out := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return out
// }


// slow linear approach using bucket sort
// this is only nice if the hand[i] is smaller at worst case
// but in this case hand[i] could be 10^9; therefore this is slow
// func isNStraightHand(hand []int, groupSize int) bool {
//     n := len(hand)
//     if n % groupSize != 0 {return false}
//     freq := map[int]int{}
//     start := math.MaxInt64
//     end := math.MinInt64
//     for i := 0; i < len(hand); i++ {
//         freq[hand[i]]++
//         start = min(start, hand[i])
//         end = max(end, hand[i])
//     }
//     rollingStartPtr := start
//     ptr := rollingStartPtr
//     currGrpSize := 0
//     lastGrpNum := math.MinInt64
//     for ptr <= end {
//         currNum := ptr
//         count := freq[currNum]
//         if count == 0 {
//             if currGrpSize > 0 {
//                 return false
//             } else {
//                 ptr++
//                 continue
//             }
//         }
//         if currGrpSize > 0 {
//             // we have a grp in-progress
//             if lastGrpNum+1 != currNum {return false}
//             lastGrpNum = currNum
//             currGrpSize++
            
//         } else {
//             // we do not have a grp in-progress
//             // i.e this is a new grp
//             lastGrpNum = currNum
//             currGrpSize = 1
//         }
//         freq[currNum]--
//         if freq[currNum] == 0 {delete(freq, currNum)}
//         if count == 1 {rollingStartPtr++}

//         // we made a grp
//         if currGrpSize == groupSize {
//             ptr = rollingStartPtr
//             currGrpSize = 0
//             lastGrpNum = math.MinInt64
//             continue
//         }
//         ptr++
//     }
//     // we must have used all the cards
//     return len(freq) == 0
// }