func romanToInt(s string) int {
    m := map[string]int {
        "I": 1,
        "IV": 4,
        "V": 5,
        "IX": 9,
        "X": 10,
        "XL": 40,
        "L": 50,
        "XC": 90,
        "C": 100,
        "CD": 400,
        "D": 500,
        "CM": 900,
        "M": 1000,
    }
    n := len(s)
    var ans int
    for i := n; i >= 1; i-- {
        if i >= 2 && m[s[i-2:i]] != 0 {
            ans += m[s[i-2:i]]
            i--
        } else {
            ans += m[s[i-1:i]]
        }
    }
    return ans
}
