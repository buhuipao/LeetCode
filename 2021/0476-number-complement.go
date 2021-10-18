func findComplement(num int) int {
    highBit := 0
    for i := 1; i <= 30; i++ {
        if num < 1<<i {
            break
        }
        highBit = i
    }
    mask := 1<<(highBit+1) - 1
    return num ^ mask
}
