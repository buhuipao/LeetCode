// 维护两个长度相差小于1的堆，一个为大根堆，一个为小根堆
type MedianFinder struct {
    lHeap []int // 大根堆，存放较小的值
    mHeap []int // 小根堆，存放较大的值
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
    m := len(this.lHeap)
    n := len(this.mHeap)
    if m == n {
        if m == 0 {
            this.lHeap = append(this.lHeap, num)
            return
        }
        if num > this.mHeap[0] {
            this.mHeap[0], num = num, this.mHeap[0]
            this.down(&(this.mHeap), -1)
        }
        this.lHeap = append(this.lHeap, num)
        this.up(&(this.lHeap), 1)
    } else {
        if m == 1 {
            if this.lHeap[0] > num {
                this.lHeap[0], num = num, this.lHeap[0]
            }
            this.mHeap = append(this.mHeap, num)
            return
        }
        if num < this.lHeap[0] {
            num, this.lHeap[0] = this.lHeap[0], num
            this.down(&(this.lHeap), 1)
        }
        this.mHeap = append(this.mHeap, num)
        this.up(&(this.mHeap), -1)
    }
}


func (this *MedianFinder) FindMedian() float64 {
    m := len(this.lHeap)
    n := len(this.mHeap)
    if m == n {
        return (float64(this.lHeap[0]) + float64(this.mHeap[0])) / 2
    } else {
        return float64(this.lHeap[0])
    }
}

func (this *MedianFinder) up(nums *[]int, t int) {
    s := len(*nums) - 1
    for s >= 1 {
        f := (s - 1) / 2
        if ((*nums)[s] - (*nums)[f]) * t > 0 {
            (*nums)[s], (*nums)[f] = (*nums)[f], (*nums)[s]
            s = f
        } else {
            break
        }
    }
}

func (this *MedianFinder) down(nums *[]int, t int) {
    n := len(*nums)
    for f := 0; f < n; {
        s := f * 2 + 1
        s1 := s + 1
        if s < n {
            if s1 < n && ((*nums)[s1] - (*nums)[s]) * t > 0 {
                s = s1
            }
            if ((*nums)[s] - (*nums)[f]) * t > 0 {
                (*nums)[f], (*nums)[s] = (*nums)[s], (*nums)[f]
                f = s
            } else {
                break
            }
        } else {
            break
        }
    }
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
