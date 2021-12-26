package hamming

import (
	"fmt"
	"math/rand"
)

type Grid [4][4]bool

func New() *Grid {
	var grid Grid
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			grid[i][j] = rand.Intn(2) == 0
		}
	}

	colOne := grid.getColumn(1)
	colTwo := grid.getColumn(3)

	grid[0][1] = getParity(append(colOne[1:], colTwo[:]...))

	colOne = grid.getColumn(2)
	colTwo = grid.getColumn(3)

	grid[0][2] = getParity(append(colOne[1:], colTwo[:]...))

	rowOne := grid.getRow(1)
	rowTwo := grid.getRow(3)

	grid[1][0] = getParity(append(rowOne[1:], rowTwo[:]...))

	rowOne = grid.getRow(2)
	rowTwo = grid.getRow(3)

	grid[2][0] = getParity(append(rowOne[1:], rowTwo[:]...))

	return &grid
}

func (g *Grid) String() string {
	var s string
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s += fmt.Sprintf("%v", g[i][j])
		}
		s += "\n"
	}
	return s
}

func getParity(bits []bool) bool {
	trues := 0

	for i := 0; i < len(bits); i++ {
		if bits[i] {
			trues++
		}
	}

	return trues%2 == 0
}

func (g *Grid) getColumn(col int) [4]bool {
	var column [4]bool
	for r, row := range g {
		column[r] = row[col]
	}
	return column
}

func (g *Grid) getRow(row int) [4]bool {
	return g[row]
}

func (g *Grid) checkOne() bool {
	colOne := g.getColumn(1)
	colTwo := g.getColumn(3)

	return getParity(append(colOne[1:], colTwo[:]...)) == colOne[0]
}

func (g *Grid) checkTwo() bool {
	colOne := g.getColumn(2)
	colTwo := g.getColumn(3)

	return getParity(append(colOne[1:], colTwo[:]...)) == colOne[0]
}

func (g *Grid) checkThree() bool {
	rowOne := g.getRow(1)
	rowTwo := g.getRow(3)

	return getParity(append(rowOne[1:], rowTwo[:]...)) == rowOne[0]
}

func (g *Grid) checkFour() bool {
	rowOne := g.getRow(2)
	rowTwo := g.getRow(3)

	return getParity(append(rowOne[1:], rowTwo[:]...)) == rowOne[0]
}

func (g *Grid) FindCorruption() (x, y int) {
	possible := [4][4]bool{
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
	}

	checkOne := g.checkOne()
	checkTwo := g.checkTwo()
	checkThree := g.checkThree()
	checkFour := g.checkFour()

	for i := 0; i < 4; i++ {
		if checkOne {
			possible[i][1], possible[i][3] = false, false
		} else {
			possible[i][0], possible[i][2] = false, false
		}

		if checkTwo {
			possible[i][2], possible[i][3] = false, false
		} else {
			possible[i][0], possible[i][1] = false, false
		}

		if checkThree {
			possible[1][i], possible[3][i] = false, false
		} else {
			possible[0][i], possible[2][i] = false, false
		}

		if checkFour {
			possible[2][i], possible[3][i] = false, false
		} else {
			possible[0][i], possible[1][i] = false, false
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if possible[i][j] {
				return i, j
			}
		}
	}

	return 0, 0
}

func (g *Grid) IsCorrupt() bool {
	x, y := g.FindCorruption()
	return x != 0 && y != 0
}

func (g *Grid) FixCorruption() {
	x, y := g.FindCorruption()
	g[x][y] = !g[x][y]
}

func (g *Grid) Check() bool {
	return g.checkOne() && g.checkTwo() && g.checkThree() && g.checkFour()
}

func (g *Grid) Corrupt() (x, y int) {
	x = rand.Intn(4)
	y = rand.Intn(4)
	for x+y == 0 {
		x = rand.Intn(4)
		y = rand.Intn(4)
	}
	g[x][y] = !g[x][y]

	return
}
