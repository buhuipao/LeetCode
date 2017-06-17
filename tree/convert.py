# _*_ coding: utf-8 _*_

# 非递归方式三种遍历树


class Tree(object):
    def __init__(self, value):
        self.val = value
        self.left = None
        self.right = None


class TreeToSequence(object):
    def convert(self, root):
        _array = [[]] * 3
        self.pre(root, _array[0])
        self.mid(root, _array[1])
        self.post(root, _array[2])

    def pre(self, root, _array):
        if not root:
            return
        _stack = []
        _stack.append(root)
        while _stack:
            # 先弹出root，然后先压入右边，再压入左边，和先序遍历顺序相反
            # 弹出后压入的左儿子，再重复上面步骤
            cur_node = _stack.pop(-1)
            _array.append(cur_node.val)
            if cur_node.right:
                _stack.append(cur_node.right)
            if cur_node.left:
                _stack.append(cur_node.left)

    def mid(self, root, _array):
        if not root:
            return
        _stack = []
        _stack.append(root)
        _cur = root.left
        # 栈不为空或者有右孩子都继续, 因为当树根结点出栈后栈为空，
        # 但是树根结点有右孩子需要继续
        while _stack or _cur:
            # 不断押入左儿子，当没有左儿子了，开始进行弹出（中序遍历）
            while _cur:
                _stack.append(_cur)
                _cur = _cur.left
            node = _stack.pop(-1)
            _array.append(node.val)
            _cur = node.right

    def post(self, root, _array):
        if not root:
            return
        _stack1, _stack2 = [], []
        _stack1.append(root)
        while _stack1:
            # 先压入根到栈1, 并弹出到栈2，压入左右到栈1
            # 之后再经过一次压栈，弹栈即为后序的左右根
            # 因为后序是左右根，所以按根右左入栈2，再从栈2弹出即可
            _cur = _stack1.pop(-1)
            _stack2.append(_cur)
            if _cur.left:
                _stack1.append(_cur.left)
            if _cur.right:
                _stack1.append(_cur.right)
        while _stack2:
            _array.append(_stack2.pop(-1).val)
