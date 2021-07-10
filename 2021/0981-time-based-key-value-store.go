type TimeMap struct {
    data map[string][]*Value
}

type Value struct {
    d string // 值
    t int   // 时间
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
    return TimeMap{data: make(map[string][]*Value)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.data[key]; !ok {
        this.data[key] = make([]*Value, 0)
    }

    this.data[key] = append(this.data[key], &Value{d: value, t: timestamp})
}


// 二分
func (this *TimeMap) Get(key string, timestamp int) string {
    if _, ok := this.data[key]; !ok {
        return ""
    }

    // 二分查找
    var ans string
    l, r := 0, len(this.data[key])-1
    for l <= r {
        m := l + (r - l)/2
        if this.data[key][m].t <= timestamp {
            ans = this.data[key][m].d
            l = m + 1
        } else {
            r = m - 1
        }
    }
    
    return ans
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
