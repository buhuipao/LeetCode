/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 import (
    "bytes"
    "fmt"
    "strings"
    "strconv"
 )

type Codec struct {
    data bytes.Buffer 
}

func Constructor() Codec {
    return Codec{data: bytes.Buffer{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    preSerialize(root, &this.data)
    return this.data.String()
}

func preSerialize(root *TreeNode, data *bytes.Buffer) {
    if root == nil {
        data.WriteString("null,")
        return
    }

    data.WriteString(fmt.Sprintf("%d,", root.Val))
    preSerialize(root.Left, data)
    preSerialize(root.Right, data)
} 

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if len(data) == 0 {
        return nil
    }

    nodes := strings.Split(data[:len(data)-1], ",")
    l, r := 0, len(nodes)-2
    return preDeserialize(nodes, &l, &r)
}

func preDeserialize(strs []string, l, r *int) *TreeNode {
    if *l > *r {
        return nil
    }

    nodeS := strs[*l]
    *l++
    if nodeS == "null" {
        return nil
    }

    val, _ := strconv.Atoi(nodeS)
    root := &TreeNode{Val: val}
    root.Left = preDeserialize(strs, l, r) 
    root.Right = preDeserialize(strs, l, r)

    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
