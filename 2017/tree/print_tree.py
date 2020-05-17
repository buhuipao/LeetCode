# _*_ coding: utf-8 _*_

# 按层打印二叉树


class TreeNode(object):
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None


class PrintTree(object):
    def printTree(self, root):
        if not root:
            return
        '''
        定义next_last为下一层的最后一个，cur_last为当前层最后一个
        temp用于存放当前行的值，resutl存放最终的结果
        '''
        next_last = cur_last = root
        _queue = [root]
        result, temp = [], []
        while _queue:
            # 在按层遍历的基础上，不断把下层最右边儿子赋值给next_last
            _cur = _queue.pop(0)
            temp.append(_cur.val)
            if _cur.left:
                _queue.append(_cur.left)
                next_last = _cur.left
            if _cur.right:
                _queue.append(_cur.right)
                next_last = _cur.right
            # 如果当前节点为此层最后的节点时，
            # 进行下层最后一个节点的赋值(cur_last=next_last)，然后才由_queue.pop(0)进入下层
            if _cur == cur_last:
                result.append(temp)
                temp = []
                cur_last = next_last
        return result
