func plusOne(digits []int) []int {
    v := 1

    for i := len(digits)-1; i >= 0; i-- {
        tmp := digits[i] + v 
        if tmp < 10 {
            digits[i] = tmp
            return digits
        }
        digits[i] = 0
        v = 1
    } 

    if v == 1 {
        return append([]int{1}, digits...)
    }

    return digits
}
