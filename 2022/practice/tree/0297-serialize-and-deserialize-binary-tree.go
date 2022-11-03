/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    data bytes.Buffer
}

func Constructor() Codec {
    return Codec{data: bytes.Buffer{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    preSialize(root, &this.data)
    return this.data.String()
}

func preSialize(root *TreeNode, data *bytes.Buffer) {
   if root == nil {
       data.WriteString("null,")
       return
   } 

   data.WriteString(fmt.Sprintf("%d,", root.Val))
   preSialize(root.Left, data)
   preSialize(root.Right, data)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    nodes := strings.Split(data[:len(data)-1], ",")

    l, r := 0, len(nodes)-2
    return preDeserialize(nodes, &l, &r)
}

func preDeserialize(nodes []string, l, r *int) *TreeNode {
    if *l > *r {
        return nil
    }

    rootNode := nodes[*l]
    *l++
    if rootNode == "null" {
        return nil
    }

    val, _ := strconv.Atoi(rootNode)
    left := preDeserialize(nodes, l, r)
    right := preDeserialize(nodes, l, r)
    return &TreeNode{
        Val: val,
        Left: left,
        Right: right,
    } 
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
