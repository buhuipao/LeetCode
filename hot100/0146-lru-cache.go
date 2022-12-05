type LinkNode struct {
    Key int
    Val int
    Pre *LinkNode
    Next *LinkNode
}

type Deque struct{
    head *LinkNode
    tail *LinkNode
}

func (d *Deque) DeleteTail() *LinkNode {
    // 无节点
    if d.tail == nil {
        return nil
    }

    tail := d.tail
    // 只有一个节点
    if d.tail.Pre == nil {
        d.tail, d.head = nil, nil
        return tail 
    }

    // 普通删除
    d.tail, d.tail.Pre = d.tail.Pre, nil

    return tail
}

func (d *Deque) Delete(node *LinkNode) {
    // 无节点
    if d.head == nil {
        return
    }

    // 仅有一个节点
    if d.head == d.tail && node == d.head {
        d.head, d.tail = nil, nil
        return
    }

    if node == d.head { // 删除头
        d.head, d.head.Next.Pre = d.head.Next, d.head.Pre
        return
    }
    if node == d.tail { // 删除尾巴
        d.tail, d.tail.Pre.Next = d.tail.Pre, d.tail.Next
        return
    }
    
    // 普通删除
    node.Pre.Next, node.Next.Pre = node.Next, node.Pre
}

func (d *Deque) AddToHead(node *LinkNode) {
    // 无节点
    if d.head == nil {
        d.head, d.tail = node, node
        return
    }

    d.head, d.head.Pre, node.Next = node, node, d.head
}

/*********************************************************/

type LRUCache struct {
    hash map[int]*LinkNode
    data *Deque
    cap int
}


func Constructor(capacity int) LRUCache {
    o := LRUCache{
        hash: make(map[int]*LinkNode),
        data: &Deque{},
        cap: capacity,
    }

    return o
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.hash[key]; !ok {
        return -1
    }

    node := this.hash[key]
    this.data.Delete(node)
    this.data.AddToHead(node)
    return node.Val
}


func (this *LRUCache) Put(key int, value int)  {
    // 新节点
    if _, ok := this.hash[key]; !ok {
        if this.cap <= len(this.hash) { // 容量不够
            tail := this.data.DeleteTail()
            if tail != nil {
                delete(this.hash, tail.Key)
            }
        }

        node := &LinkNode{Key: key, Val: value}
        this.hash[key] = node
        this.data.AddToHead(node)
        return
    }

    // 旧节点
    this.hash[key].Val = value
    // 一次读取，让其被放置到Deque头
    this.Get(key)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
