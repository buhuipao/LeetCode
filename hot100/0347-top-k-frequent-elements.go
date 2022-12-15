// hash + heap
// 1）生成频率统计
// 2）维护小根堆
func topKFrequent(nums []int, k int) []int {
    if k == 0 {
        return nil
    }

    m := make(map[int]int)
    for i := range nums {
        m[nums[i]]++
    }

    // 生成小根堆
    heap := make([]int, 0, k)
    for key := range m {
        if len(heap) < k {
            heap = append(heap, key)
            son := len(heap) - 1
            // 上浮调整堆
            for son >= 1 {
                father := (son - 1) / 2
                // 如果儿子不小于比父亲值，则终止进行上浮
                if m[heap[son]] >= m[heap[father]] {
                    break
                }

                heap[son], heap[father] = heap[father], heap[son]
                son = father
            }
        } else {
            // 替换出堆顶的最小值，塞入较大值
            // 进行下沉调整
            if m[key] > m[heap[0]] {
                heap[0] = key
                father := 0
                for father < k {
                    son1, son2 := father*2+1, father*2+2
                    if son1 >= k {
                        break
                    }

                    // 找出两个儿子中更小的替换下沉的父亲
                    if son2 < k && m[heap[son2]] < m[heap[son1]] {
                        son1 = son2
                    }

                    if m[heap[son1]] >= m[heap[father]] {
                        break
                    }

                    heap[son1], heap[father] = heap[father], heap[son1]
                    father = son1
                }
            }
        }
    }

    return heap
}
