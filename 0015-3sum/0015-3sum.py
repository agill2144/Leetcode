class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        nums.sort()
        tuples = []
        for i in range(n-2):
            if (i > 0 and nums[i] == nums[i-1]) :
                continue
            low = i+1
            high = n-1
            while low < high:
                sum = nums[i] + nums[low] + nums[high]
                if sum == 0:
                    pair = [nums[i], nums[low], nums[high]]
                    tuples.append(pair)
                    low+=1
                    high-=1
                    while(low < high and nums[low] == nums[low-1]):
                        low+=1
                    while(low < high and nums[high] == nums[high+1]):
                        high-=1
                elif sum > 0:
                    high-=1
                else:
                    low+=1
        return tuples

