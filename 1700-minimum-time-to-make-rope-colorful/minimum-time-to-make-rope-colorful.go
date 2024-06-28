import ("math")
func minCost(colors string, neededTime []int) int {
    count := 0
    min := math.MaxInt64
    second := math.MaxInt64
    totalTime := 0
    for i := 0; i < len(colors); i++ {
        if i == 0 {
            min = neededTime[i]
            count = 1
            continue
        }

        if colors[i] == colors[i-1] {
            count++
            if neededTime[i] < min {
                second = min
                min = neededTime[i]
            } else if neededTime[i] < second {
                second = neededTime[i]
            }

            if count == 2 {
                count = 1
                totalTime += min
                min = second
                second = math.MaxInt64
            }
        } else {
            count = 1
            min = neededTime[i]
            second = math.MaxInt64
        }
    }
    return totalTime
}

// simulate using a stack
// time = o(n)
// space = o(n)
// func minCost(colors string, neededTime []int) int {
//     type stNode struct {
//         color byte
//         time int
//     }
//     st := []stNode{}
//     totalTime := 0
//     for i := 0; i < len(colors); i++ {
//         currColor := colors[i]
//         currTime := neededTime[i]
//         pushCurr := true
//         for len(st) != 0 && currColor == st[len(st)-1].color {
//             // drop the one that takes least amount of time
//             top := st[len(st)-1]
//             topTime := top.time
//             if currTime < topTime {
//                 pushCurr = false
//                 totalTime += currTime
//                 break
//             } else {
//                 st = st[:len(st)-1]
//                 totalTime += topTime
//             }
//         }
//         if pushCurr {
//             st = append(st, stNode{color: currColor, time: currTime})
//         }
//     }
//     return totalTime
// }