/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
type NestedIterator struct {
    list []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{list: nestedList}
}

func (it *NestedIterator) Next() int {
    v := it.list[0].GetInteger()
    it.list = it.list[1:]
    return v
}

func (it *NestedIterator) HasNext() bool {
    // 不停循环，直到列表为空或者第一个元素为整数值
    var head *NestedInteger
    var nodes []*NestedInteger
    for len(it.list) != 0 && !it.list[0].IsInteger() {
        head, it.list = it.list[0], it.list[1:]
        nodes = head.GetList()
        for i := len(nodes)-1; i >= 0; i-- {
            it.list = append([]*NestedInteger{nodes[i]}, it.list...)
        }
    }

    return len(it.list) != 0
}
