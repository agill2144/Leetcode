


// sort and pick cheapest ones first
// time: o(nlogn) + o(min(costs, coins))
// space: o(n) if merge sort since it may use extra temp array to merge
func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    i := 0
    count := 0
    for coins > 0 && i < len(costs) && costs[i] <= coins {
        coins -= costs[i]
        count++
        i++
    }
    return count
}