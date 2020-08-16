func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if len(image) == 0 || image[sr][sc] == newColor {
        return image
    }
    rawValue := image[sr][sc]
    dfs(&image, sr, sc, rawValue, newColor)
    return image
}

func dfs(image *[][]int, row, col int, oldV, newV int) {
    (*image)[row][col] = newV
    nexts := [][]int{
        []int{0, -1},
        []int{0, 1},
        []int{-1, 0},
        []int{1, 0},
    }
    for _, v := range nexts {
        newR, newC := row + v[0], col + v[1]
        if newR >= 0 && newR < len(*image) && newC >= 0 && newC < len((*image)[0]) && (*image)[newR][newC] == oldV {
            dfs(image, newR, newC, oldV, newV)
        }
    }

}
