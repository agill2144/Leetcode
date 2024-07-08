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
func trap(height []int) int {
    n := len(height)
    left := make([]int, n)
    right := make([]int, n)
    for i := 0; i < n; i++ {
        if i == 0 || i == n-1 {continue}
        maxLeftSoFar := max(left[i-1], height[i-1])
        maxRightSoFar := max(right[n-1-i+1], height[n-1-i+1])
        left[i] = maxLeftSoFar
        right[n-1-i] = maxRightSoFar
    }
    total := 0
    for i := 0; i < n; i++ {
        leftH, rightH := left[i],right[i]
        canBeTrappedAboveIthBuilding := min(leftH, rightH)-height[i]
        if canBeTrappedAboveIthBuilding > 0 {
            total += canBeTrappedAboveIthBuilding
        }
    }
    return total
}