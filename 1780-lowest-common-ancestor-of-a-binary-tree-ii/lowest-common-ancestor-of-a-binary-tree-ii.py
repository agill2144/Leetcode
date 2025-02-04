# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        p_found = False
        q_found = False
        def dfs(r: 'TreeNode') -> 'TreeNode':
            nonlocal p_found
            nonlocal q_found
            # base
            
            if not r:
                return None

            # logic
            left = dfs(r.left)
            right = dfs(r.right)
            if r is p:
                p_found = True
                return r
            if r is q:
                q_found = True
                return r

            if left and right:
                return r
            if left:
                return left
            return right
        res = dfs(root)
        if p_found and q_found:
            return res
        return None