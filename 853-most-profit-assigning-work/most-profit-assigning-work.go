/*
    approach: brute force
    - combine the 2 arrays, diff and profit
    - for each worker ( each worker has a skill )
    - find out which highest paying job this worker can have
        - i.e a job whose difficulty <= current worker skill
        - 1 job at max per worker
    - keep a total profit

    approach: optimizing above
    - "find out which highest paying job this worker can have"
        - if these jobs were in a heap ( sorted by profit ), then we can have the above answer
    - sort the combined array by difficulty
    - so for each worker , loop over the merged/combined array
        - as long as the current job diff <= current worker skill 
        - toss this job into a max heap (sorted by profit )
        - as soon as a job diff > current worker skill, we can break
            - the merged array is combined and its safe to break because next jobs will require higher skill
            - i.e the merged array is sorted in asc order by diff
        - take the highest paying job and add its profit to current global total
        - DO NOT POP AND DO NOT REMOVE FROM HEAP
        - because same job can be done by all workers, therefore we want to persist all the jobs
        - now for next worker, take its skill, continue moving the ptr for combined array from where it stopped last time
    
    n = len(workers)
    m = len(merged array)    
    time = o(m) + o(mlogm) + o(n * mlogm)
    space = o(m)


    approach: greedy
    - combine diff and profit into a single merged array
    - sort the merged array by difficulty
    - sort the workers array
    - keep track of global profit
    - *keep track of last profit gained by worker*
    - loop over workers 
        - ( remember this arr is sorted in asc order )
        - ( remember that merged arr is also sorted  )
    - for each worker;
        - keep updating lastProfit ( max lastProfit, currentProfit ) 
        - as long as job skill <= curr worker skill
        - once job skill > current worker skill, break
        - take the lastProfit and add to global total
        - * do not rest lastProfit *
            - because if current worker can earn $lastProfit
            - and workers are sorted
            - it means next worker skill >= current worker skill ( because sorted arr )
            - this means, next worker can definetely take the last job too ( a job can be done by multiple workers )
            - thats why we do not reset and thats why we do not have start ALL OVER AGAIN on the combined/merged array
            - we kept the lastProfit and lastProfit will continue to work for next worker too (unless next worker finds a better job)
*/
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    n := len(difficulty)
    merged := [][]int{}
    for i := 0; i < n; i++ {
        merged = append(merged, []int{difficulty[i], profit[i]})
    }
    sort.Slice(merged, func(i, j int)bool{
        if merged[i][0] == merged[j][0] {
            return merged[i][1] < merged[j][1]
        }
        return merged[i][0] < merged[j][0]        
    })

    sort.Ints(worker)
    totalProfit := 0
    workerProfit := 0
    ptr := 0
    for i := 0; i < len(worker); i++ {
        workerSkillLevel := worker[i]
        for ptr < len(merged) {
            if merged[ptr][0] <= workerSkillLevel {
                workerProfit = max(workerProfit, merged[ptr][1])
                ptr++
            } else {
                break
            }
        }
        totalProfit += workerProfit
    }
    return totalProfit
}