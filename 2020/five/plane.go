package five

import "strings"

func ParseTicket(s string) int {
	// Get Row
	min := 0
	max := 127
	middle := min + ((max - min) / 2)
	r := -1
	for i, c := range s {
		switch c {
		case 'F':
			if i == 6 {
				r = min
			}
			max = middle
		case 'B':
			if i == 6 {
				r = max
			}
			min = middle + 1
		default:
			break
		}

		middle = min + ((max - min) / 2)

	}

	col := -1

	min = 0
	max = 7
	middle = min + ((max - min) / 2)
	// Get col
	for i, c := range s[len(s)-3:] {
		switch c {
		case 'L':
			if i == 2 {
				col = min
			}
			max = middle
		case 'R':
			if i == 2 {
				col = max
			}
			min = middle + 1
		default:
			break
		}
		middle = min + ((max - min) / 2)

	}

	//r * 8 + col
	return (r * 8) + col
}

func HighestTicket(s string) int {
	tickets := strings.Split(s, "\n")

	m := -1
	for _, t := range tickets {
		clean := strings.TrimSpace(t)
		if clean == "" {
			continue
		}
		id := ParseTicket(clean)

		if id > m {
			m = id
		}
	}

	return m
}
