/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

// 字典树
 type TrieNode struct {
    children [26]*TrieNode
    val      int
}

type MapSum struct {
    root *TrieNode
    cnt  map[string]int
}

func Constructor() MapSum {
    return MapSum{&TrieNode{}, map[string]int{}}
}

func (m *MapSum) Insert(key string, val int) {
    delta := val
    if m.cnt[key] > 0 {
        delta -= m.cnt[key]
    }
    m.cnt[key] = val
    node := m.root
    for _, ch := range key {
        ch -= 'a'
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{}
        }
        node = node.children[ch]
        node.val += delta
    }
}

func (m *MapSum) Sum(prefix string) int {
    node := m.root
    for _, ch := range prefix {
        ch -= 'a'
        if node.children[ch] == nil {
            return 0
        }
        node = node.children[ch]
    }
    return node.val
}
