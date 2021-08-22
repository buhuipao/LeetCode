func escapeGhosts(ghosts [][]int, target []int) bool {
    source := []int{0, 0}
    distance := manhattanDistance(source, target)
    for _, ghost := range ghosts {
        if manhattanDistance(ghost, target) <= distance {
            return false
        }
    }
    return true
}

func manhattanDistance(point1, point2 []int) int {
    return abs(point1[0]-point2[0]) + abs(point1[1]-point2[1])
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
