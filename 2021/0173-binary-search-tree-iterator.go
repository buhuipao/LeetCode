/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    min int
    cur *TreeNode
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    stack := []*TreeNode{}
    for root != nil {
        stack = append(stack, root)
        root = root.Left
    }
    return BSTIterator{stack: stack, cur: root, min: ^int(^uint(0)>>1)}
}


func (this *BSTIterator) Next() int {
    if !this.HasNext() {
        return this.min
    }
    if this.cur != nil {
        for this.cur != nil {
            this.stack = append(this.stack, this.cur)
            this.cur = this.cur.Left
        }
    }
    node := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    this.cur = node.Right
    return node.Val
}


func (this *BSTIterator) HasNext() bool {
    return this.cur != nil || len(this.stack) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
