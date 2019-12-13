package eleven

import "testing"

import "github.com/stretchr/testify/assert"

func Test_Robot(t *testing.T) {
	p := []int{3, 8, 1005, 8, 326, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 29, 2, 1003, 17, 10, 1006, 0, 22, 2, 106, 5, 10, 1006, 0, 87, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 65, 2, 7, 20, 10, 2, 9, 17, 10, 2, 6, 16, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 99, 1006, 0, 69, 1006, 0, 40, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 127, 1006, 0, 51, 2, 102, 17, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1002, 8, 1, 155, 1006, 0, 42, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 180, 1, 106, 4, 10, 2, 1103, 0, 10, 1006, 0, 14, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 213, 1, 1009, 0, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 239, 1006, 0, 5, 2, 108, 5, 10, 2, 1104, 7, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 272, 2, 1104, 12, 10, 1, 1109, 10, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 302, 1006, 0, 35, 101, 1, 9, 9, 1007, 9, 1095, 10, 1005, 10, 15, 99, 109, 648, 104, 0, 104, 1, 21102, 937268449940, 1, 1, 21102, 1, 343, 0, 1105, 1, 447, 21101, 387365315480, 0, 1, 21102, 1, 354, 0, 1105, 1, 447, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21101, 0, 29220891795, 1, 21102, 1, 401, 0, 1106, 0, 447, 21101, 0, 248075283623, 1, 21102, 412, 1, 0, 1105, 1, 447, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 0, 984353760012, 1, 21102, 1, 435, 0, 1105, 1, 447, 21102, 1, 718078227200, 1, 21102, 1, 446, 0, 1105, 1, 447, 99, 109, 2, 21202, -1, 1, 1, 21102, 40, 1, 2, 21101, 0, 478, 3, 21101, 468, 0, 0, 1106, 0, 511, 109, -2, 2106, 0, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 473, 474, 489, 4, 0, 1001, 473, 1, 473, 108, 4, 473, 10, 1006, 10, 505, 1102, 1, 0, 473, 109, -2, 2105, 1, 0, 0, 109, 4, 1202, -1, 1, 510, 1207, -3, 0, 10, 1006, 10, 528, 21102, 1, 0, -3, 22102, 1, -3, 1, 22101, 0, -2, 2, 21101, 0, 1, 3, 21102, 1, 547, 0, 1105, 1, 552, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 575, 2207, -4, -2, 10, 1006, 10, 575, 21202, -4, 1, -4, 1105, 1, 643, 21202, -4, 1, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21102, 1, 594, 0, 1106, 0, 552, 22102, 1, 1, -4, 21101, 1, 0, -1, 2207, -4, -2, 10, 1006, 10, 613, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 635, 22101, 0, -1, 1, 21101, 0, 635, 0, 106, 0, 510, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}
	x, y := 50, 50

	painter := NewPainter(p, x, y)

	painter.hull[y][x].color = 1

	i := painter.Run()

	assert.Equal(t, 249, i)
	// painter.Close()

	painter.printHull()
}
