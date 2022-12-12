/*
    approach: greedy
    - We need to assign each bike to the closest worker possible
    - So for each worker;
        - we will calculate its distance from all bikes
        - distance calculation of 2 cordinates will be done using manhattan distance -- abs(y2-y1)+abs(x2-x1)
        - and store it in a map by distance { dist: [ [workerId, bikeId] ]}
    - then we will need to start with lowest possible distance first from the map we generated
        - The keys in the map will be distances and a list of workerId,bikeID representing that these 2 pairs are at this distance.
        - Inorder to start from the smallest distance ( why? because we need to assign a bike to its closest worker - so it makes sense to start with pairs with smallest distance first) 
            - we will maintain and min var as soon as we have a distance smaller than min, than we will update our min to new min dist
    - Remember we need to assign each worker to its closest bike
        - so our output array will be representing workers assigned to which bike
        - idx is the workerId, and the value at a specific idx means - this workerIdx is assigned this bikeId
        - However we need to keep track of which all bikes and workers are still available..
        - We could maintain 2 lists of bools.
            - 1. len of workers - representing available workers
                - once a worker is assigned, we will change that workerIdx to true - so no longer avail
            - 2. len of bikes - representing available bikes
                - once a bike is occupied, we will change that bikeIdx to true, so no longer avail
            - Or flip the script and maintain 2 hashsets
                - 1. workerSet : if a workerId exists in set that means its no longer available
                - 2. bikeSet : if a bikeId exists in set that means its no longer available
    - Keep in mind that there may be more bikes than workers
    - So if our result array == len(workers) == this means we have no more workers avail to assign remaining bikes to - return and exit with whatever we have
    time: o(mn)
    space: o(mn)
*/

func assignBikes(workers [][]int, bikes [][]int) []int {
    distMap := map[int][][]int{}
    min := math.MaxInt64
    max := math.MinInt64
    m := len(workers)
    n := len(bikes)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dist := calcDist(workers[i], bikes[j])
            distMap[dist] = append(distMap[dist], []int{i,j})
            if dist < min {min=dist}
            if dist > max {max=dist}
        }
    }
    
    result := make([]int, m)
    assignedWorkers := make([]bool, m)
    occupiedBikes := make([]bool, n)
    count := len(workers)
    for i := min; i <= max; i++ {
        wbPairs := distMap[i]
        for _, wbPair := range wbPairs {
            wr := wbPair[0]
            bk := wbPair[1]
            if !assignedWorkers[wr] && !occupiedBikes[bk] {
                assignedWorkers[wr] = true
                occupiedBikes[bk] = true
                result[wr] = bk
                count--
            }
            if count == 0 {
                return result
            }
        }
    }
    return result
}

func calcDist(wr []int, bk []int) int {
    wx := wr[0]
    wy := wr[1]
    bx := bk[0]
    by := bk[1]
    return abs(wy-by)+abs(wx-bx)
}

func abs(n int) int {
    if n < 0 {
        return n*-1
    }
    return n
}