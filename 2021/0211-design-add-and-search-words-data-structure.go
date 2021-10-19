// 字典树
 type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}

func (t *TrieNode) Insert(word string) {
    node := t
    for _, ch := range word {
        ch -= 'a'
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

type WordDictionary struct {
    trieRoot *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{&TrieNode{}}
}

func (d *WordDictionary) AddWord(word string) {
    d.trieRoot.Insert(word)
}

func (d *WordDictionary) Search(word string) bool {
    var dfs func(int, *TrieNode) bool
    dfs = func(index int, node *TrieNode) bool {
        if index == len(word) {
            return node.isEnd
        }
        ch := word[index]
        if ch != '.' {
            child := node.children[ch-'a']
            if child != nil && dfs(index+1, child) {
                return true
            }
        } else {
            for i := range node.children {
                child := node.children[i]
                if child != nil && dfs(index+1, child) {
                    return true
                }
            }
        }
        return false
    }
    return dfs(0, d.trieRoot)
}
