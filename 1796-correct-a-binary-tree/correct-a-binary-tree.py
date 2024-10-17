# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def correctBinaryTree(self, root: TreeNode) -> TreeNode:
        if not root:
            return None

        q = [[root, root]] # <node, parent>
        while q:
            q_size = len(q)
            ## contiains nodes visited from this level so far
            processed_level_nodes = set()
            for _ in range(q_size):
                dq = q.pop(0)
                curr_node = dq[0]
                parent = dq[1]
                if curr_node.right:
                    if curr_node.right in processed_level_nodes:
                        if parent.left and parent.left == curr_node:
                            parent.left = None
                        elif parent.right and parent.right == curr_node:
                            parent.right = None
                        return root
                    q.append([curr_node.right, curr_node])
                if curr_node.left:
                    q.append([curr_node.left, curr_node])
                processed_level_nodes.add(curr_node)
        return root
