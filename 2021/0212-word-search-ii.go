type Trie struct {
    children [26]*Trie
    word     string
}

func (t *Trie) Insert(word string) {
    node := t
    for _, ch := range word {
        ch -= 'a'
        if node.children[ch] == nil {
            node.children[ch] = &Trie{}
        }
        node = node.children[ch]
    }
    node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
    t := &Trie{}
    for _, word := range words {
        t.Insert(word)
    }

    m, n := len(board), len(board[0])
    seen := map[string]bool{}

    var dfs func(node *Trie, x, y int)
    dfs = func(node *Trie, x, y int) {
        ch := board[x][y]
        node = node.children[ch-'a']
        if node == nil {
            return
        }

        if node.word != "" {
            seen[node.word] = true
        }

        board[x][y] = '#'
        for _, d := range dirs {
            nx, ny := x+d.x, y+d.y
            if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
                dfs(node, nx, ny)
            }
        }
        board[x][y] = ch
    }
    for i, row := range board {
        for j := range row {
            dfs(t, i, j)
        }
    }

    ans := make([]string, 0, len(seen))
    for s := range seen {
        ans = append(ans, s)
    }
    return ans
}

