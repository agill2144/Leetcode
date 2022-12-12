class Solution:
    def canJump(self, nums: List[int]) -> bool:
        memo = {}
        def dfs(startIdx: int) -> bool:
            ## base
            if startIdx >= len(nums)-1:
                return True
            
            ## logic
            for i in range(nums[startIdx],0,-1):
                newIdx = i + startIdx
                if newIdx in memo:
                    if memo.get(newIdx) == True:
                        return True
                else:
                    res = dfs(newIdx)
                    memo[newIdx] =  res
                    if res:
                        return True
            
            return False
        return dfs(0)