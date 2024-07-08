// we have to compute how much water can be trapped on top of each building
// we need some sort of gurantee that there exists a maxLeft and maxRight without extra space
// we can use 2 vars to keep track of maxLeft (leftWall) and maxRight (rightWall)
// and use 2 ptrs (left and right) to process either left building or right building
// if we have a leftWall <= rightWall; meaning rightWall is greater than or equal to leftWall
// meaning we have a maxRight, then can probably process left building
// because 1. we have a leftWall(maxLeft) and 2. rightWall(rightMax) and 3. we have left ptr( building to process)
// time = o(n)
// space = o(1)
// func trap(height []int) int {
//     n := len(height)
//     left := 0
//     right := n-1
//     leftWall := height[left]
//     rightWall := height[right]
//     total := 0
//     for left < right {
//         // if leftWall <= rightWall
//         // then we know for-sure that there is a rightWall >= leftWall
//         // this means, we can for sure figure out how much can be trapped on left building
//         // therefore process left building
//         // ( its like we have a maxLeft and a maxRight from non-optimal solution )
//         if leftWall <= rightWall {
//             // we can calcuate how much we can trap on top of left building
//             // we can only trap water on left building IF left building's height < leftWall
//             if height[left] <= leftWall {
//                 total += (leftWall-height[left])
//                 // we have computed and saved how much we can trap on top of left building while keeping leftWall as-is
//                 // therefore we can evaluate the next left building
//                 // and therefore we can move away from left building
//                 left++
//             } else {
//                 // we have a new leftWall since height[left] >= leftWall
//                 leftWall = height[left]
//             }

//         // otherwise leftWall > rightWall
//         // which means, we for-sure have a leftWall > rightWall
//         // which means, we can figure out how much water can be trapped on top of right building
//         // ( its like we have a maxLeft and a maxRight from non-optimal solution )
//         } else {
//             // we can calculate how much we can trap on top of right building
//             // we can only trap water on right building if right building's height < rightWall
//             if height[right] <= rightWall {
//                 total += (rightWall-height[right])
//                 // we have computed and saved how much we can trap on top of right building while keeping rightWall as-is
//                 // therefore we can evaluate the next right building
//                 // and therefore we can move away from right building
//                 right--
//             } else {
//                 // we have a new righWall since right building's height > rightWall
//                 rightWall = height[right]
//             }
//         }
//     }
//     return total
// }

/*
    approach: prefix / suffix
    - for each building;
    - find how much can be trapped above a building
    - we can trap water above a building if;
        - there exists a left building > curr building height
        - there exists a right buidling > curr building height
            __          __
            ||    __    ||
            || __ || __ ||
        - we can trap water above building whose height is 0
            - can trap 2 units of water ; min(leftBuilding, rightBuilding)-currHeight
        - we can trap water above building whose height is 1
            - can trap 1 unit of water ; min(leftBuilding, rightBuilding)-currHeight
        - we can trap water above building whose height is 0
            - can trap 2 units of water ; min(leftBuilding, rightBuilding)-currHeight
        - leftBuilding and rightBuilding are the max heights ( max vals ) excluding current building
        - i.e max val from left and max val from right
        - i.e max prefix from left and max suffix from right
    - once left and right prefixes
    - we can figure out how much can be trapped above each ith building
        - total += min(left[i], right[i]) - height[i]

    time = o(2n) = o(n)
    space = o(2n) = o(n)
*/
// func trap(height []int) int {
//     n := len(height)
//     left := make([]int, n)
//     right := make([]int, n)
//     for i := 0; i < n; i++ {
//         if i == 0 || i == n-1 {continue}
//         maxLeftSoFar := max(left[i-1], height[i-1])
//         maxRightSoFar := max(right[n-1-i+1], height[n-1-i+1])
//         left[i] = maxLeftSoFar
//         right[n-1-i] = maxRightSoFar
//     }
//     total := 0
//     for i := 0; i < n; i++ {
//         leftH, rightH := left[i],right[i]
//         canBeTrappedAboveIthBuilding := min(leftH, rightH)-height[i]
//         if canBeTrappedAboveIthBuilding > 0 {
//             total += canBeTrappedAboveIthBuilding
//         }
//     }
//     return total
// }


// brute force
func trap(height []int) int {
    n := len(height)
    total := 0
    for i := 1; i < n-1; i++ {
        curr := height[i]
        left := 0
        for k := 0; k < i; k++ {
            left = max(left, height[k])
        }
        right := 0
        for k := i+1; k < n; k++ {
            right = max(right, height[k])
        }
        wall := min(left, right)
        if wall > curr {
            total += wall-curr            
        }
    }
    return total
}