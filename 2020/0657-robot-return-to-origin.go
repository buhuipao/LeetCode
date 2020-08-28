func judgeCircle(moves string) bool {
    m := map[rune]int{
        'U': -1,
        'D': 1,
        'R': -1,
        'L': 1,
    }
    a, b := 0, 0
    count := map[rune]*int {
        'U': &a,
        'D': &a,
        'R': &b,
        'L': &b,
    }
    for _, c := range moves {
        *count[c] += m[c]
    }
    return a == 0 && b == 0
}
