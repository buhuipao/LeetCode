# _*_ coding: utf-8 _*_
# 树的先序，中序，后序递归遍历


class Tree(object):
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

    def convert(self, root):
        _array = [[]] * 3
        self.xian(self, root, _array[0])
        self.zhong(self, root, _array[1])
        self.hou(self, root, _array[2])

    def pre(self, root, _array):
        if not root:
            return 0
        else:
            _array.append(root.value)
            self.xian(self, root.left, _array)
            self.xian(self, root.right, _array)

    def mid(self, root, _array):
        if not root:
            return 0
        else:
            self.zhong(self, self.left, _array)
            _array.append(root.value)
            self.zhong(self, self.right, _array)

    def post(self, root, _array):
        if not root:
            return 0
        else:
            self.hou(self, self.left, _array)
            self.hou(self, self.right, _array)
            _array.append(root.value)
