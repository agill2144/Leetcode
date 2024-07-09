// we have to compute how much water can be trapped on top of each building
// we need some sort of gurantee that there exists a maxLeft and maxRight without extra space
// we can use 2 vars to keep track of maxLeft (leftWall) and maxRight (rightWall)
// and use 2 ptrs (left and right) to process either left building or right building
// if we have a leftWall <= rightWall; meaning rightWall is greater than or equal to leftWall
// meaning we have a maxRight, then can probably process left building
// because 1. we have a leftWall(maxLeft) and 2. rightWall(rightMax) and 3. we have left ptr( building to process)
// time = o(n)
// space = o(1)
func trap(height []int) int {
    n := len(height)
    left := 0
    right := n-1
    leftWall := height[left]
    rightWall := height[right]
    total := 0
    for left < right {
        if leftWall <= rightWall {
            if leftWall >= height[left] {
                total += (leftWall-height[left])
                left++
            } else {
                leftWall = height[left]
            }
        } else {
            if rightWall >= height[right] {
                total += (rightWall-height[right])
                right--
            }  else {
                rightWall = height[right]
            }
        }
    }
    return total
}

/*

    - in brute force, we were looking for maxLeft and maxRight for each building
    - we can pre-compute this in prefix and suffix array
    - and then processing the heights array in linear time

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

    time: o(n) + o(n)
    spae = o(n)
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
// time: o(n^2)
// space = o(1)
// func trap(height []int) int {
//     n := len(height)
//     total := 0
//     for i := 1; i < n-1; i++ {
//         curr := height[i]
//         left := 0
//         for k := 0; k < i; k++ {
//             left = max(left, height[k])
//         }
//         right := 0
//         for k := i+1; k < n; k++ {
//             right = max(right, height[k])
//         }
//         wall := min(left, right)
//         if wall > curr {
//             total += wall-curr            
//         }
//     }
//     return total
// }