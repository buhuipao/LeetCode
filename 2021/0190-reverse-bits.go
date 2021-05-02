/*
// 一个左移一个右移
func reverseBits(num uint32) uint32 {
    var ans uint32
    for i := 0; i < 32; i++ {
        ans <<= 1
        ans += (num & 1)
        num >>= 1
    }
    return ans
}
*/

func reverseBits(num uint32) uint32 {
    flags := make([]bool, 32)
    for i := 0; i < 32; i++ {
        if (num & (1 << i)) != 0 {
            flags[i] = true
        }
    }
    ans := uint32(0)
    for i, v := range flags {
        if v {
            ans += 1 << (31-i)
        }
    }
    return ans
}
