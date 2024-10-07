# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        pFound = False
        qFound = False
        def dfs(r):
            nonlocal pFound
            nonlocal qFound
            # base
            if not r:
                return None

            # logic
            left = dfs(r.left)
            right = dfs(r.right)

            if r == p:
                pFound = True
                return r
            if r == q:
                qFound = True
                return r

            if left and right:
                return r
            if left:
                return left
            return right
        res = dfs(root)
        if pFound and qFound:
            return res
        return None