func countSubTrees(n int, edges [][]int, labels string) []int {
    m := make(map[int][]int)
    for _, edge := range edges {
        m[edge[0]] = append(m[edge[0]], edge[1])
        m[edge[1]] = append(m[edge[1]], edge[0])
    }

    result := make([]int, n)
    dfs(0, -1, m, labels, &result)
    return result
}

func dfs(node int, parent int, m map[int][]int, labels string, result *[]int) map[byte]int {
    counter := make(map[byte]int)
    counter[labels[node]]++

    for _, n := range m[node] {
        if n == parent {
            continue
        }
        cnt := dfs(n, node, m, labels, result)
        for k, v := range cnt {
            counter[k] += v
        }
    }
    (*result)[node] = counter[labels[node]]
    return counter
}