type NumMatrix struct {
    data [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    o := NumMatrix{
        data: make([][]int, len(matrix)+1),
    }
    for i := range o.data {
        o.data[i] = make([]int, len(matrix[0])+1)
    }

    var sum int
    for i := range matrix {
        sum = 0
        for j := range matrix[i] {
            sum += matrix[i][j]
            o.data[i+1][j+1] = sum + o.data[i][j+1]
        }
    }

    return o
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.data[row2+1][col2+1] - this.data[row1][col2+1] - this.data[row2+1][col1] + this.data[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
