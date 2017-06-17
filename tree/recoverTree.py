# _*_ coding: utf-8 _*_


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def recoverTree(self, root):
        """
        :type root: TreeNode
        :rtype: void Do not return anything, modify root in-place instead.
        只需要中序遍历，假如不是单调递增的，就交换即可
        """
        if not root:
            return

        _stack = [root]
        first_enode = None
        second_enode = None
        temp = None
        _cur = root.left
        while _cur or _stack:
            while _cur:
                _stack.append(_cur)
                _cur = _cur.left
            node = _stack.pop(-1)
            # 如果不是单调递增就保存此节点为first_enode
            # 当得到第一个错误点之后，再得到第二个错误节点，最后交换value
            if first_enode:
                if int(temp.val) < int(first_enode_next.val) or int(first_enode.val) < int(node.val):
                    second_enode = temp
                    sencod_enode.val, first_enode.val = first_enode.val, second_enode.val
                    break
            if not temp or int(temp.val) <= int(node.val):
                temp = node
            else:
                first_enode = temp
                first_enode_next = node
            _cur = node.right
