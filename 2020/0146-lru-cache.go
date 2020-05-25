type LRUCache struct {
    hash map[int]*LinkNode
    data *Deque
    cap int
}

type Deque struct {
    Head *LinkNode
    Tail *LinkNode
}

type LinkNode struct {
    Key int
    Value int
    Pre *LinkNode
    Next *LinkNode
}

func (d *Deque) Delete(node *LinkNode) {
    // 只有一个节点
    if d.Head == d.Tail && node == d.Head {
        d.Head, d.Tail = nil, nil
        return
    }
    if node == d.Head { // 删除头
        d.Head, d.Head.Next.Pre = d.Head.Next, nil
    } else if node == d.Tail { // 删除尾巴
        d.Tail, d.Tail.Pre.Next = d.Tail.Pre, nil
    } else {
        node.Next.Pre, node.Pre.Next = node.Pre, node.Next
    }
}

func (d *Deque) AddToHead(node *LinkNode) {
    if d.Head == nil {
        d.Head, d.Tail = node, node
    } else {
       d.Head, node.Next, d.Head.Pre = node, d.Head, node
    }
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        hash: make(map[int]*LinkNode, 0),
        data: &Deque{},
        cap: capacity,
    }
    return lru
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.hash[key]; !ok {
        return -1
    }
    node := this.hash[key]
    this.data.Delete(node)
    this.data.AddToHead(node)
    return this.hash[key].Value
}


func (this *LRUCache) Put(key int, value int)  {
    // 新节点
    if _, ok := this.hash[key]; !ok {
        if this.cap <= len(this.hash) { // 容量不够
            delete(this.hash, this.data.Tail.Key)
            this.data.Delete(this.data.Tail)
        }
        node := &LinkNode{Key: key, Value: value}    
        this.hash[key] = node
        this.data.AddToHead(node)
        return
    }
    // 旧节点
    this.hash[key].Value = value
    this.Get(key)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
