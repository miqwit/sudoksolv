package main

import (
	"fmt"
	"log"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Contains the full grid, with secured numbers
var grid [9][9]int

// Contains a grid of options for each empty cell.
// If a cell is not empty, slice of option is empty.
var gridOptions [9][9][]int

// printGrid will display to the standard output a nice ASCII
// version of the 2-dimensional array representing the sudoku grid
func printGrid(withHints bool) {
	fmt.Println("+---+---+---+---+---+---+---+---+---+")
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Print("| ")
			if (grid[row][col] != 0) {
				fmt.Print(grid[row][col])
			} else {
				if (withHints && len(gridOptions[row][col]) == 1) {
					fmt.Print("\033[31m◆\033[0m")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print(" ")
		}
		fmt.Println("|")
		fmt.Println("+---+---+---+---+---+---+---+---+---+")
	}
}

func printGridOptions() {
	fmt.Println("+---------------+---------------+---------------+---------------+---------------+---------------+---------------+---------------+---------------+")
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Print("| ")
			if (grid[row][col] != 0) {
				fmt.Printf("\033[31m%-13d\033[0m", grid[row][col])
			} else {
				var strOptions = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(gridOptions[row][col])), " "), "[]")
				fmt.Printf("%-13s", strOptions)
			}
			fmt.Print(" ")
		}
		fmt.Println("|")
		fmt.Println("+---------------+---------------+---------------+---------------+---------------+---------------+---------------+---------------+---------------+")
	}
}

// strToGrid converts a string to a Sudoku grid. The string must
// contain only digits from 0 (empty cell) to 9. The string will fill
// the grid line by line. For example, the string
//   120000050800400030000050958...
// will fill the grid
//   +---+---+---+---+---+---+---+---+---+
//   | 1 | 2 |   |   |   |   |   | 5 |   |
//   +---+---+---+---+---+---+---+---+---+
//   | 8 |   |   | 4 |   |   |   | 3 |   |
//   +---+---+---+---+---+---+---+---+---+
//   ...
func strToGrid(str string) {
	// check string is 81 values
	if (len(str) != 81) {
		log.Fatal(errors.New("Not a valid grid. Submit 81 values."))
	}

	// check all values are valid
	var validGrid = regexp.MustCompile(`[0-9]{81}`)
	if (!validGrid.MatchString(str)) {
		log.Fatal(errors.New("Not a valid grid. Values must be numbers from 0 to 9."))
	}

	// convert string to grid
	var row, col int = 0, 0
	for _, ch := range str {
		grid[row][col], _ = strconv.Atoi(string(ch))
		col++
		if (col == 9) {
			col = 0
			row++
		}
	}
}

// isInRow returns true if given value is in given row
func isInRow(row int, val int) bool {
	for col := 0; col < 9; col++ {
		if (grid[row][col] == val) {
			return true
		}
	}
	return false
}

// isInCol returns true if given value is in given col
func isInCol(col int, val int) bool {
	for row := 0; row < 9; row++ {
		if (grid[row][col] == val) {
			return true
		}
	}
	return false
}

// getSquareFromRowCol returns the number of the square given
// the column and row. Squares are distributed as following
// 3x3 subgrids:
// +---+---+---+---+---+---+---+---+---+
// |   |   |   |   |   |   |   |   |   |
// +---+---+---+---+---+---+---+---+---+
// |   | 1 |   |   | 2 |   |   | 3 |   |
// +---+---+---+---+---+---+---+---+---+
// |   |   |   |   |   |   |   |   |   |
// +---+---+---+---+---+---+---+---+---+
// |   |   |   |   |   |   |   |   |   |
// +---+---+---+---+---+---+---+---+---+
// |   | 4 |   |   | 5 |   |   | 6 |   |
// +---+---+---+---+---+---+---+---+---+
// |   |   |   |   |   |   |   |   |   |
// +---+---+---+---+---+---+---+---+---+
// |   |   |   |   |   |   |   |   |   |
// +---+---+---+---+---+---+---+---+---+
// |   | 7 |   |   | 8 |   |   | 9 |   |
// +---+---+---+---+---+---+---+---+---+
// |   |   |   |   |   |   |   |   |   |
// +---+---+---+---+---+---+---+---+---+
func getSquareFromRowCol(row int, col int) int {
	var rowOffset int = (row / 3) * 3
	var colOffset int = (col / 3)

	return (colOffset + 1) + rowOffset
}

// isInSquare return true is the given value is already present
// in the given square. Squares are numbered from 1 to 9, 
// see getSquareFromRowCol documentation.
func isInSquare(square int, val int) bool {
	var cols [3]int
	var rows [3]int

	var rowOffset int = ((square - 1) / 3) * 3
	rows = [3]int{0 + rowOffset, 1 + rowOffset, 2 + rowOffset}

	var colOffset int = ((square - 1) % 3) * 3
	cols = [3]int{0 + colOffset, 1 + colOffset, 2 + colOffset}

	for _, row := range rows {
		for _, col := range cols {
			if (grid[row][col] == val) {
				return true
			}
		}
	}
	return false
}

// countEmptyCells returns the number of zeros in the grid.
func countEmptyCells() int {
	var numEmpty int = 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if (grid[row][col] == 0) {
				numEmpty++
			}
		}
	}

	return numEmpty
}

// For each empty cell in the grid, list the possible options
func listOptionsPerEmptyCell() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if (grid[row][col] != 0) {
				continue
			}

			// list each number and add it as an option if
			// not in row, line or square already
			// fmt.Printf("Working on row %d col %d\n", row, col)
			var options []int
			for value := 1; value < 10; value++ {
				if (isInRow(row, value)) {
					continue
				}
				
				if (isInCol(col, value)) {
					continue
				}
				
				if (isInSquare(getSquareFromRowCol(row, col), value)) {
					continue
				}

				options = append(options, value)
			}
			gridOptions[row][col] = options
			if (len(options) == 1) {
				fmt.Printf("r%d,c%d: \033[31m%v\033[0m\n", row+1, col+1, options)
			} else {
				// fmt.Printf("r%d,c%d: %v\n", row+1, col+1, options)
			}
		}
	}
}

// fillSecuredOptions will replace in grid what gridOptions found
// as the only reliable option.
func fillSecuredOptions() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if (len(gridOptions[row][col]) == 1) {
				grid[row][col] = gridOptions[row][col][0]
				gridOptions[row][col] = []int{} // reset options for this cell.
			}
		}
	}
}

func reduceOptionsFromUniqueOccurenceGeneric(rowMin int, rowMax int, colMin int, colMax int, zoneType string) {
	var dict = make(map[int]int)

	for row := rowMin; row <= rowMax; row++ {
		for col := colMin; col <= colMax; col++ {
			for _, option := range gridOptions[row][col] {
				dict[option] = dict[option]+1
			}
		}
	}

	// If an option has only one possibility in the zone, set it as the only option.
	var valueToFix int
	for option, amount := range dict {
		if (amount == 1) {
			fmt.Printf("In %s, value %d can only be in one place\n", zoneType, option)
			valueToFix = option
		}
	}

	// Browse again this zone, and force this value when present.
	for row := rowMin; row <= rowMax; row++ {
		for col := colMin; col <= colMax; col++ {
			for _, option := range gridOptions[row][col] {
				if (option == valueToFix) {
					gridOptions[row][col] = []int{valueToFix}
					continue
				}
			}
		}
	}

}

// reduceOptionsFromUniqueOccurence will select options that can't be elsewhere
// on the square, the row or the column. The given cell can have multiple options
// but only one cell of the squar/row/column can ultimately host it; e.g. the other
// cells does not have this possible option.
func reduceOptionsFromUniqueOccurence() {
	// Browse all squares
	for square := 1; square <= 9; square++ {
		var rowOffset int = ((square - 1) / 3) * 3
		var colOffset int = ((square - 1) % 3) * 3
		
		reduceOptionsFromUniqueOccurenceGeneric(0 + rowOffset, 2 + rowOffset, 0 + colOffset, 2 + colOffset, fmt.Sprintf("square %d", square))
	}

	// Browse all rows
	for row := 0; row < 9; row++ {
		reduceOptionsFromUniqueOccurenceGeneric(row, row, 0, 8, fmt.Sprintf("row %d", row + 1))
	}

	// Browse all cols
	for col := 0; col < 9; col++ {
		reduceOptionsFromUniqueOccurenceGeneric(0, 8, col, col, fmt.Sprintf("col %d", col + 1))
	}

	// printGridOptions()
}

func main() {
	// level 3
	// strToGrid("120000050800400030000050948013200000400503007000001820731080000040006009060000084")
	// level 3-4
	// strToGrid("100030002903040600200000300000308700010207030006904000001000009004070501600080003")
	// strToGrid("090000000183090000065001700000170200010208090004035000006700340000010586000000020")
	// level 4
	// strToGrid("480006007300002490000004020000300281000000000731005000090700000043500009100600053")
	strToGrid("006000300435009007701600000870002010000000000060900082000006105900100276007000800")
	
	printGrid(false)

	var remains int = countEmptyCells()
	listOptionsPerEmptyCell() // fills gridOptions

	for (remains > 0) {
		reduceOptionsFromUniqueOccurence()
		printGrid(true)
		fillSecuredOptions()
		listOptionsPerEmptyCell()
		printGrid(true)

		if (countEmptyCells() == remains) {
			break
		}

		remains = countEmptyCells()
	}

	if (remains != 0) {
		log.Fatal(errors.New("Could not solve."))
	}
}