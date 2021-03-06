package ten

import "testing"

import "github.com/stretchr/testify/assert"

func Test_Astroid_BearingToAstroid(t *testing.T) {
	testCases := []struct {
		x1, y1 int
		x2, y2 int
		expect float64
	}{
		{
			x1:     1,
			y1:     1,
			x2:     2,
			y2:     2,
			expect: 315,
		},
		{
			x1:     1,
			y1:     1,
			x2:     3,
			y2:     3,
			expect: 315,
		},
		{
			x1:     2,
			y1:     2,
			x2:     1,
			y2:     1,
			expect: 135,
		},
		{
			x1:     1,
			y1:     2,
			x2:     1,
			y2:     1,
			expect: 180,
		},
		{
			x1:     1,
			y1:     1,
			x2:     1,
			y2:     2,
			expect: 0,
		},
		{
			x1:     3,
			y1:     4,
			x2:     2,
			y2:     2,
			expect: 153.43494882292202,
		},
		{
			x1:     3,
			y1:     4,
			x2:     1,
			y2:     0,
			expect: 153.43494882292202,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			a := &Astroid{
				X: tc.x1,
				Y: tc.y1,
			}

			b := &Astroid{
				X: tc.x2,
				Y: tc.y2,
			}

			theta := a.BearingToAstroid(b)
			assert.Equal(t, tc.expect, theta)
		})
	}

}

func Test_Astroid_CheckIfLineOfSite(t *testing.T) {
	testCases := []struct {
		x, y          int
		in            string
		expectedCount int
	}{
		{
			x:             3,
			y:             4,
			in:            ".#..#\n.....\n#####\n....#\n...##",
			expectedCount: 8,
		},
		{
			x:             5,
			y:             8,
			in:            "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####",
			expectedCount: 33,
		},
		{
			x:             1,
			y:             2,
			in:            "#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###.",
			expectedCount: 35,
		},
		{
			x:             6,
			y:             3,
			in:            ".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#..",
			expectedCount: 41,
		},
		{
			x:             11,
			y:             13,
			in:            ".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##",
			expectedCount: 210,
		},
		{
			x:             26,
			y:             29,
			in:            ".............#..#.#......##........#..#\n.#...##....#........##.#......#......#.\n..#.#.#...#...#...##.#...#.............\n.....##.................#.....##..#.#.#\n......##...#.##......#..#.......#......\n......#.....#....#.#..#..##....#.......\n...................##.#..#.....#.....#.\n#.....#.##.....#...##....#####....#.#..\n..#.#..........#..##.......#.#...#....#\n...#.#..#...#......#..........###.#....\n##..##...#.#.......##....#.#..#...##...\n..........#.#....#.#.#......#.....#....\n....#.........#..#..##..#.##........#..\n........#......###..............#.#....\n...##.#...#.#.#......#........#........\n......##.#.....#.#.....#..#.....#.#....\n..#....#.###..#...##.#..##............#\n...##..#...#.##.#.#....#.#.....#...#..#\n......#............#.##..#..#....##....\n.#.#.......#..#...###...........#.#.##.\n........##........#.#...#.#......##....\n.#.#........#......#..........#....#...\n...............#...#........##..#.#....\n.#......#....#.......#..#......#.......\n.....#...#.#...#...#..###......#.##....\n.#...#..##................##.#.........\n..###...#.......#.##.#....#....#....#.#\n...#..#.......###.............##.#.....\n#..##....###.......##........#..#...#.#\n.#......#...#...#.##......#..#.........\n#...#.....#......#..##.............#...\n...###.........###.###.#.....###.#.#...\n#......#......#.#..#....#..#.....##.#..\n.##....#.....#...#.##..#.#..##.......#.\n..#........#.......##.##....#......#...\n##............#....#.#.....#...........\n........###.............##...#........#\n#.........#.....#..##.#.#.#..#....#....\n..............##.#.#.#...........#.....",
			expectedCount: 299,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			a := NewAstroid(tc.x, tc.y)

			astroids := GetAstroids(tc.in)

			a.CheckIfLineOfSite(astroids)
			a.Sort()
			assert.Equal(t, tc.expectedCount, a.LineOfSiteCount())
		})
	}
}

func Test_Astroid_ReturnNthAstroidToDestroy(t *testing.T) {
	testCases := []struct {
		x, y      int
		in        string
		eX, eY, n int
	}{
		{
			x:  8,
			y:  3,
			in: ".#....#####...#..\n##...##.#####..##\n##...#...#.#####.\n..#.....#...###..\n..#.#.....#....##",
			eX: 8,
			eY: 1,
			n:  1,
		},
		// {
		// 	x:             5,
		// 	y:             8,
		// 	in:            "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####",
		// 	expectedCount: 33,
		// },
		// {
		// 	x:             1,
		// 	y:             2,
		// 	in:            "#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###.",
		// 	expectedCount: 35,
		// },
		// {
		// 	x:             6,
		// 	y:             3,
		// 	in:            ".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#..",
		// 	expectedCount: 41,
		// },
		{
			x:  11,
			y:  13,
			in: ".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##",
			eX: 8,
			eY: 2,
			n:  200,
		},
		{
			x:  26,
			y:  29,
			in: ".............#..#.#......##........#..#\n.#...##....#........##.#......#......#.\n..#.#.#...#...#...##.#...#.............\n.....##.................#.....##..#.#.#\n......##...#.##......#..#.......#......\n......#.....#....#.#..#..##....#.......\n...................##.#..#.....#.....#.\n#.....#.##.....#...##....#####....#.#..\n..#.#..........#..##.......#.#...#....#\n...#.#..#...#......#..........###.#....\n##..##...#.#.......##....#.#..#...##...\n..........#.#....#.#.#......#.....#....\n....#.........#..#..##..#.##........#..\n........#......###..............#.#....\n...##.#...#.#.#......#........#........\n......##.#.....#.#.....#..#.....#.#....\n..#....#.###..#...##.#..##............#\n...##..#...#.##.#.#....#.#.....#...#..#\n......#............#.##..#..#....##....\n.#.#.......#..#...###...........#.#.##.\n........##........#.#...#.#......##....\n.#.#........#......#..........#....#...\n...............#...#........##..#.#....\n.#......#....#.......#..#......#.......\n.....#...#.#...#...#..###......#.##....\n.#...#..##................##.#.........\n..###...#.......#.##.#....#....#....#.#\n...#..#.......###.............##.#.....\n#..##....###.......##........#..#...#.#\n.#......#...#...#.##......#..#.........\n#...#.....#......#..##.............#...\n...###.........###.###.#.....###.#.#...\n#......#......#.#..#....#..#.....##.#..\n.##....#.....#...#.##..#.#..##.......#.\n..#........#.......##.##....#......#...\n##............#....#.#.....#...........\n........###.............##...#........#\n#.........#.....#..##.#.#.#..#....#....\n..............##.#.#.#...........#.....",
			n:  200,
			eX: 14,
			eY: 19,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			a := NewAstroid(tc.x, tc.y)

			astroids := GetAstroids(tc.in)

			a.CheckIfLineOfSite(astroids)
			n := a.ReturnNthAstroidToDestroy(tc.n)

			assert.Equal(t, tc.eX, n.X)
			assert.Equal(t, tc.eY, n.Y)

		})
	}
}

func Test_FindAstroidWithMaxSight(t *testing.T) {
	testCases := []struct {
		x, y int
		in   string
	}{
		{
			x:  3,
			y:  4,
			in: ".#..#\n.....\n#####\n....#\n...##",
		},
		{
			x:  5,
			y:  8,
			in: "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####",
		},
		{
			x:  1,
			y:  2,
			in: "#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###.",
		},
		{
			x:  6,
			y:  3,
			in: ".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#..",
		},
		{
			x:  11,
			y:  13,
			in: ".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##",
		},
		{
			x:  26,
			y:  29,
			in: ".............#..#.#......##........#..#\n.#...##....#........##.#......#......#.\n..#.#.#...#...#...##.#...#.............\n.....##.................#.....##..#.#.#\n......##...#.##......#..#.......#......\n......#.....#....#.#..#..##....#.......\n...................##.#..#.....#.....#.\n#.....#.##.....#...##....#####....#.#..\n..#.#..........#..##.......#.#...#....#\n...#.#..#...#......#..........###.#....\n##..##...#.#.......##....#.#..#...##...\n..........#.#....#.#.#......#.....#....\n....#.........#..#..##..#.##........#..\n........#......###..............#.#....\n...##.#...#.#.#......#........#........\n......##.#.....#.#.....#..#.....#.#....\n..#....#.###..#...##.#..##............#\n...##..#...#.##.#.#....#.#.....#...#..#\n......#............#.##..#..#....##....\n.#.#.......#..#...###...........#.#.##.\n........##........#.#...#.#......##....\n.#.#........#......#..........#....#...\n...............#...#........##..#.#....\n.#......#....#.......#..#......#.......\n.....#...#.#...#...#..###......#.##....\n.#...#..##................##.#.........\n..###...#.......#.##.#....#....#....#.#\n...#..#.......###.............##.#.....\n#..##....###.......##........#..#...#.#\n.#......#...#...#.##......#..#.........\n#...#.....#......#..##.............#...\n...###.........###.###.#.....###.#.#...\n#......#......#.#..#....#..#.....##.#..\n.##....#.....#...#.##..#.#..##.......#.\n..#........#.......##.##....#......#...\n##............#....#.#.....#...........\n........###.............##...#........#\n#.........#.....#..##.#.#.#..#....#....\n..............##.#.#.#...........#.....",
		},
	}

	for i, tc := range testCases {
		if i != 5 {
			continue
		}
		t.Run("", func(t *testing.T) {
			astroids := GetAstroids(tc.in)

			a := FindAstroidWithMaxSight(astroids)

			assert.Equal(t, tc.x, a.X)
			assert.Equal(t, tc.y, a.Y)

		})
	}
}

func Test_GetAstroids(t *testing.T) {
	testCases := []struct {
		in    string
		count int
	}{
		{
			in:    ".#..#\n.....\n#####\n....#\n...##",
			count: 10,
		},
		{
			in:    "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####",
			count: 40,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			a := GetAstroids(tc.in)

			assert.Len(t, a, tc.count)
		})
	}
}
