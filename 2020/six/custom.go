package six

import (
	"strings"
)

func CountGroupQuestions(s string) int {
	m := make(map[string]int)

	people := strings.Split(s, "\n")
	peopleCount := 0
	for _, p := range people {
		p := strings.TrimSpace(p)
		if p == "" {
			continue
		}
		peopleCount++
		for _, a := range p {
			m[string(a)]++
		}
	}

	c := 0
	for _, v := range m {
		if v == peopleCount {
			c++
		}
	}
	return c
}

func CountAllGroupsQuestions(s string) []int {
	lines := strings.Split(s, "\n")
	counts := make([]int, 0)
	b := strings.Builder{}

	for _, l := range lines {
		clean := strings.TrimSpace(l)
		if clean == "" {
			sub := b.String()
			if sub == "" {
				continue
			}
			i := CountGroupQuestions(sub)
			b.Reset()
			counts = append(counts, i)
			continue
		}

		b.WriteString(clean)
		b.WriteString("\n")
	}

	// grab the last input
	sub := b.String()
	if sub != "" {
		i := CountGroupQuestions(b.String())
		counts = append(counts, i)
	}

	return counts
}

func Count(s string) int {
	counts := CountAllGroupsQuestions(s)

	count := 0

	for _, c := range counts {
		count += c
	}

	return count
}
