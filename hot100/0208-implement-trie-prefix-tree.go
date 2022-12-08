type Trie struct {
    children [26]*Trie
    isEnd bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    node := this
    var index int
    for _, c := range word {
        index = int(c - 'a')
        if node.children[index] ==  nil {
            node.children[index] = &Trie{}
        }
        // 选择下级节点
        node = node.children[index]
    }

    node.isEnd = true
}


func (this *Trie) Search(word string) bool {
    node := this.searchNode(word)
    return node != nil && node.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    return this.searchNode(prefix) != nil
}

func (this *Trie) searchNode(prefix string) *Trie {
    node := this
    var index int
    for _, c := range prefix {
        index = int(c - 'a')
        if node.children[index] == nil {
            return nil
        }
        // 选择下级节点
        node = node.children[index]
    }

    return node
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
