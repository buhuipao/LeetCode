// 大根堆实现，用堆保存k个最小值，第k大是堆顶
func kthSmallest(matrix [][]int, k int) int {
    heap := make([]int, 0, k)
    n := len(matrix)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            pushHeap(&heap, matrix[i][j], k)
        }
    }
    return heap[0]
}


func pushHeap(nums *[]int, x int, k int) {
    if len(*nums) < k {
        *nums = append(*nums, x)
        // 上浮操作
        for s := len(*nums) - 1; s >= 1; {
            f := (s - 1) / 2
            if (*nums)[s] > (*nums)[f] {
                (*nums)[s], (*nums)[f] = (*nums)[f], (*nums)[s]
                s = f
            } else {
                break
            }
        }
    } else {
        // 发现更小的值
        if x < (*nums)[0] {
            (*nums)[0] = x
            // 下沉操作，注意不要越界
            for f := 0; f < k; {
                s := f * 2 + 1
                s1 := s + 1
                if s1 < k {
                    if (*nums)[s1] > (*nums)[s] {
                        s = s1
                    }
                    if (*nums)[s] > (*nums)[f] {
                        (*nums)[s], (*nums)[f]= (*nums)[f], (*nums)[s]
                    } else{
                        break
                    }
                } else if s < k {
                    if (*nums)[s] > (*nums)[f] {
                        (*nums)[s], (*nums)[f]= (*nums)[f], (*nums)[s]
                    } else{
                        break
                    }
                } else {
                    break
                }
                f = s
            }
        }
    }
}
