# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        # 先序遍历二叉树
        if root is None:
            return ''
        stack = [root]
        values = []
        while stack:
            node = stack.pop()
            if node is not None:
                values.append(str(node.val))
                stack.extend([node.right, node.left])
            else:
                values.append('null')
        # 删除多余的空节点
        while values[-1] == 'null':
            values.pop()
        return ','.join(values)
        
    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if not data:
            return None
        values = data.split(',')
        root = TreeNode(int(values[0]))
        stack, node = [root], root
        # 左孩子标志
        # 如果被设立新节点应该是一个左孩子，否则为右孩子
        as_left = True
        for v in values[1:]:
            # 先序遍历中一定是先保存空的左孩子，接着保存空的右孩子，所以遇到空值时需要改变标志位
            # 需要弹栈，并且改变as_left标志
            # 虽然我写这么详细，但是半年后我一定不会记得自己是怎么想出来的 -_-""
            if v == 'null':
                node = stack.pop()
                as_left = False
            else:
                temp = TreeNode(int(v))
                stack.append(temp)
                if as_left:
                  node.left = temp
                else:
                    node.right = temp
                node = temp
                # 节点值不为空时，下一个一定是左孩子，所以需要改变标志位
                as_left = True

        return root
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
