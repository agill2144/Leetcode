func trap(height []int) int {
    n := len(height)
    left := 0
    right := n-1
    leftWall := height[left]
    rightWall := height[right]
    total := 0
    for left < right {
        if leftWall <= rightWall {
            // we can calcuate how much we can trap on top of left building
            // we can only trap water on left building IF left building's height < leftWall
            if height[left] <= leftWall {
                total += (leftWall-height[left])
                // we have computed and saved how much we can trap on top of left building while keeping leftWall as-is
                // therefore we can evaluate the next left building
                left++
            } else {
                // we have a new leftWall since height[left] >= leftWall
                leftWall = height[left]
            }
        } else {
            // we can calculate how much we can trap on top of right building
            // we can only trap water on right building if right building's height < rightWall
            if height[right] <= rightWall {
                total += (rightWall-height[right])
                // we have computed and saved how much we can trap on top of right building while keeping rightWall as-is
                // therefore we can evaluate the next right building
                right--
            } else {
                // we have a new righWall since right building's height > rightWall
                rightWall = height[right]
            }
        }
    }
    return total
}

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