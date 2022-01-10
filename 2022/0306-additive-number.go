func stringAdd(x, y string) string {
    res := []byte{}
    carry, cur := 0, 0
    for x != "" || y != "" || carry != 0 {
        cur = carry
        if x != "" {
            cur += int(x[len(x)-1] - '0')
            x = x[:len(x)-1]
        }
        if y != "" {
            cur += int(y[len(y)-1] - '0')
            y = y[:len(y)-1]
        }
        carry = cur / 10
        cur %= 10
        res = append(res, byte(cur)+'0')
    }
    for i, n := 0, len(res); i < n/2; i++ {
        res[i], res[n-1-i] = res[n-1-i], res[i]
    }
    return string(res)
}

func valid(num string, secondStart, secondEnd int) bool {
    n := len(num)
    firstStart, firstEnd := 0, secondStart-1
    for secondEnd <= n-1 {
        third := stringAdd(num[firstStart:firstEnd+1], num[secondStart:secondEnd+1])
        thirdStart := secondEnd + 1
        thirdEnd := secondEnd + len(third)
        if thirdEnd >= n || num[thirdStart:thirdEnd+1] != third {
            break
        }
        if thirdEnd == n-1 {
            return true
        }
        firstStart, firstEnd = secondStart, secondEnd
        secondStart, secondEnd = thirdStart, thirdEnd
    }
    return false
}

func isAdditiveNumber(num string) bool {
    n := len(num)
    for secondStart := 1; secondStart < n-1; secondStart++ {
        if num[0] == '0' && secondStart != 1 {
            break
        }
        for secondEnd := secondStart; secondEnd < n-1; secondEnd++ {
            if num[secondStart] == '0' && secondStart != secondEnd {
                break
            }
            if valid(num, secondStart, secondEnd) {
                return true
            }
        }
    }
    return false
}
