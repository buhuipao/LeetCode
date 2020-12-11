func predictPartyVictory(senate string) string {
    var radiant, dire []int
    for i, s := range senate {
        if s == 'R' {
            radiant = append(radiant, i)
        } else {
            dire = append(dire, i)
        }
    }
    // 比较两个站队队首的序号，序号更小的可以干掉对方队伍的队首，同时把自己加到本队队尾
    for len(radiant) > 0 && len(dire) > 0 {
        if radiant[0] < dire[0] {
            radiant = append(radiant, radiant[0]+len(senate))
        } else {
            dire = append(dire, dire[0]+len(senate))
        }
        radiant = radiant[1:]
        dire = dire[1:]
    }
    if len(radiant) > 0 {
        return "Radiant"
    }
    return "Dire"
}
