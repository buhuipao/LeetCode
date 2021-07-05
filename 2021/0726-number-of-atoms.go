func countOfAtoms(formula string) string {
    i, n := 0, len(formula)

    parseAtom := func() string {
        start := i
        i++ // 扫描，跳过首字母
        for i < n && unicode.IsLower(rune(formula[i])) { 
            i++ // 扫描首字母后的小写字母
        }
        return formula[start:i]
    }

    parseNum := func() (num int) {
        if i == n || !unicode.IsDigit(rune(formula[i])) { 
            return 1 // 不是数字，视作 1
        }
        for ; i < n && unicode.IsDigit(rune(formula[i])); i++ { 
            num = num*10 + int(formula[i]-'0') // 扫描数字
        }
        return
    }

    stk := []map[string]int{{}}
    for i < n {
        if ch := formula[i]; ch == '(' {
            i++
            stk = append(stk, map[string]int{}) // 将一个空的哈希表压入栈中，准备统计括号内的原子数量
        } else if ch == ')' {
            i++
            num := parseNum() // 括号右侧数字
            atomNum := stk[len(stk)-1]
            stk = stk[:len(stk)-1] // 弹出括号内的原子数量
            for atom, v := range atomNum {
                stk[len(stk)-1][atom] += v * num // 将括号内的原子数量乘上 num，加到上一层的原子数量中
            }
        } else {
            atom := parseAtom()
            num := parseNum()
            stk[len(stk)-1][atom] += num // 统计原子数量
        }
    }

    atomNum := stk[0]
    type pair struct {
        atom string
        num  int
    }
    pairs := make([]pair, 0, len(atomNum))
    for k, v := range atomNum {
        pairs = append(pairs, pair{k, v})
    }
    sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })

    ans := []byte{}
    for _, p := range pairs {
        ans = append(ans, p.atom...)
        if p.num > 1 {
            ans = append(ans, strconv.Itoa(p.num)...)
        }
    }
    return string(ans)
}
