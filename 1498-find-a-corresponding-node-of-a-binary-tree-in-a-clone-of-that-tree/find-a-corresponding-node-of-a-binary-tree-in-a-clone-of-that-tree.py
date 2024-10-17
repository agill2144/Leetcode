# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        def dfs(r :TreeNode) -> TreeNode:
            # base
            if not r:
                return None

            # logic
            left = dfs(r.left)
            right = dfs(r.right)
            if left:
                return left
            if right:
                return right
            if r.val == target.val:
                return r
            return None
        return dfs(cloned)        