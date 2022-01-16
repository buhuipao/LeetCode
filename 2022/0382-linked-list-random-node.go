type Solution struct {
    head *ListNode
}

func Constructor(head *ListNode) Solution {
    return Solution{head}
}

func (s *Solution) GetRandom() (ans int) {
    for node, i := s.head, 1; node != nil; node = node.Next {
        if rand.Intn(i) == 0 { // 1/i 的概率选中（替换为答案）
            ans = node.Val
        }
        i++
    }
    return
}
