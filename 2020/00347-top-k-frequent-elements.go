// hash + heap
// 1）生成频率统计
// 2）维护小根堆
func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int, 0)
    for i := range nums {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = 0
        }
        m[nums[i]]++
    }
    // 生成小根堆
    heap := make([]int, 0, k)
    for key := range m {
        if len(heap) < k {
            heap = append(heap, key)
            s := len(heap) - 1
            for s >= 1 {
                f := (s - 1) / 2
                if m[heap[s]] < m[heap[f]] {
                    heap[s], heap[f] = heap[f], heap[s]
                    s = f
                } else {
                    break
                }
            }
        } else {
            if m[key] > m[heap[0]] {
                heap[0] = key
                f := 0
                for f < k {
                    // 两个儿子
                    s := f * 2 + 1
                    s1 := s + 1
                    if s < k {
                        if s1 < k && m[heap[s1]] < m[heap[s]] {
                            s = s1
                        }
                        if m[heap[s]] < m[heap[f]] {
                            heap[s], heap[f] = heap[f], heap[s]
                            f = s
                        } else {
                            break
                        }
                    } else {
                        break
                    }
                }
            }
        }
    }
    return heap
}
