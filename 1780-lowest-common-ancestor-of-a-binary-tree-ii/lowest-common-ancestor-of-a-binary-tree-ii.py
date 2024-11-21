# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if not p or not q or not root:
            return None
        p_exists = False
        q_exists = False
        def dfs(r: 'TreeNode') -> 'TreeNode':
            nonlocal p_exists
            nonlocal q_exists
            # base
            if not r:
                return None

            # logic
            left = dfs(r.left)
            right = dfs(r.right)
            if r == p:
                p_exists = True
                return r
            if r == q:
                q_exists = True
                return r
            if left and right:
                return r
            return left if left else right
        res = dfs(root)
        if p_exists and q_exists:
            return res
        return None