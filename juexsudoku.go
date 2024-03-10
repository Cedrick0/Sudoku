package main

func main() {

type SudokuGrid = [][]int
type MemoizationMap [9][9][10]bool

var rowsUsed, colsUsed, boxesUsed, boxesUsed MemoizationMap

func initializeMemoizationMaps(grid *SudokuGrid) {
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      num := grid[i][j]
      if num != 0 {
        idx := (i/3)*3 + j/3
        rowsUsed[i][j][num] = true
        colsUsed[i][j][num] = true
        boxesUsed[idx][j][num] = true
      }
    }
  }
}

func solveSudoku(grid *SudokuGrid) bool {
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      if grid[i][j] == 0 {
        for num := 1; num <= 9; num++ {
          if isValid(i, j, num) {
            placeNumber(grid, i, j, num)

            if solveSudoku(grid) {
              return true
            } else {
              removeNumber(grid, i, j, num)
  p          }
          }
        }
        return false
      }
    }
  }
  return true
}

func isValid(row, col, num int) bool {
  idx := (row / 3) * 3 + col / 3
  return !rowsUsed[row][num] && !colsUsed[col][num] && !boxesUsed[idx][num]
}

func placeNumber(grid *SudokuGrid, row, col, num int) {
  idx := (row / 3) * 3 + col / 3

  grid[row][col] = num
  rowsUsed[row][num] = true
  colsUsed[col][num] = true
  boxesUsed[idx][num] = true
}

func removeNumber(grid *SudokuGrid, row, col, num int) {
  idx := (row / 3) * 3 + col / 3

  grid[row][col] = 0
  rowsUsed[row][num] = false
  colsUsed[col][num] = false
  boxesUsed[idx][num] = false
}
}