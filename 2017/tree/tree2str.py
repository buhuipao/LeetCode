# _*_ coding: utf-8 _*_

# 先序遍历序列化树


class TreeToString(object):
    def toString(self, root):
        if not root:
            return ''
        _stack, _str = [root], ''
        # 在先序遍历的基础上，进行字符串拼接就好
        while _stack:
            _cur = _stack.pop(-1)
            if _cur:
                _str += str(_cur.val) + '!'
                _stack.append(_cur.right)
                _stack.append(_cur.left)
            else:
                _str += '#!'
        return _str
