package hamming

import (
	"fmt"
	"math/rand"
)

type Code [4][4]bool

func New() *Code {
	code := new(Code)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			code[i][j] = rand.Intn(2) == 0
		}
	}

	colOne := code.getColumn(1)
	colTwo := code.getColumn(3)

	code[0][1] = getParity(append(colOne[1:], colTwo[:]...))

	colOne = code.getColumn(2)
	colTwo = code.getColumn(3)

	code[0][2] = getParity(append(colOne[1:], colTwo[:]...))

	rowOne := code.getRow(1)
	rowTwo := code.getRow(3)

	code[1][0] = getParity(append(rowOne[1:], rowTwo[:]...))

	rowOne = code.getRow(2)
	rowTwo = code.getRow(3)

	code[2][0] = getParity(append(rowOne[1:], rowTwo[:]...))

	return code
}

func (c *Code) String() string {
	var s string
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s += fmt.Sprintf("%v", c[i][j])
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

func (c *Code) getColumn(col int) [4]bool {
	var column [4]bool
	for r, row := range c {
		column[r] = row[col]
	}
	return column
}

func (c *Code) getRow(row int) [4]bool {
	return c[row]
}

func (c *Code) checkOne() bool {
	colOne := c.getColumn(1)
	colTwo := c.getColumn(3)

	return getParity(append(colOne[1:], colTwo[:]...)) == colOne[0]
}

func (c *Code) checkTwo() bool {
	colOne := c.getColumn(2)
	colTwo := c.getColumn(3)

	return getParity(append(colOne[1:], colTwo[:]...)) == colOne[0]
}

func (c *Code) checkThree() bool {
	rowOne := c.getRow(1)
	rowTwo := c.getRow(3)

	return getParity(append(rowOne[1:], rowTwo[:]...)) == rowOne[0]
}

func (c *Code) checkFour() bool {
	rowOne := c.getRow(2)
	rowTwo := c.getRow(3)

	return getParity(append(rowOne[1:], rowTwo[:]...)) == rowOne[0]
}

func (c *Code) FindCorruption() (x, y int) {
	possible := [4][4]bool{
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
	}

	checkOne := c.checkOne()
	checkTwo := c.checkTwo()
	checkThree := c.checkThree()
	checkFour := c.checkFour()

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

func (c *Code) IsCorrupt() bool {
	x, y := c.FindCorruption()
	return x != 0 && y != 0
}

func (c *Code) FixCorruption() {
	x, y := c.FindCorruption()
	c[x][y] = !c[x][y]
}

func (c *Code) Corrupt() (x, y int) {
	x = rand.Intn(4)
	y = rand.Intn(4)
	for x+y == 0 {
		x = rand.Intn(4)
		y = rand.Intn(4)
	}
	c[x][y] = !c[x][y]

	return
}
