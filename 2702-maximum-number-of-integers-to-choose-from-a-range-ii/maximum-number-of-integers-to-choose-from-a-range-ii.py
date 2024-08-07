class Solution:
    def maxCount(self, banned, n, maxSum):
        low, high = 0, n

        while low <= high:
            mid = (low + high)//2
            total = mid*(mid+1)//2

            for i in banned:
                if mid >= i: total -= i

            if maxSum >= total: low = mid + 1
            else: high = mid - 1

        return high - sum(i <= high for i in set(banned))






        

        



