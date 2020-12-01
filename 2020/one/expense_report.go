package one

import "errors"

var (
	// ErrorInputToSmall is returned when the provided report is to small to meet the sum requirements
	ErrorInputToSmall = errors.New("input to small")
	// ErrorNoSumFound is returned when no sum can be found to meet the requirements
	ErrorNoSumFound = errors.New("no sum found")
)

// ProductOfN takes a list of ints finds n ints that add up to targetSum and then returns their product
func ProductOfN(report []int, targetSum, n int) (int, error) {
	r, err := findSum(report, targetSum, n)
	if err != nil {
		return 0, err
	}

	product := 1

	for _, p := range r {
		product = product * p
	}

	return product, nil
}

func findSum(report []int, targetSum, numberInSum int) ([]int, error) {
	if len(report) < numberInSum {
		return nil, ErrorInputToSmall
	}

	for i, v := range report {
		r := findSumRecursive(report[i+1:], targetSum, numberInSum-1, v)
		if r != nil {
			r = append(r, v)
			return r, nil
		}
	}

	return nil, ErrorNoSumFound
}

func findSumRecursive(report []int, targetSum, depth, base int) []int {
	if len(report) < depth {
		return nil
	}

	for i, v := range report {
		if depth > 1 {
			intBase := base + v
			if i+1 < len(report) {
				r := findSumRecursive(report[i+1:], targetSum, depth-1, intBase)
				if r != nil {
					r = append(r, v)
					return r
				}
			}

		} else {
			if v+base == targetSum {
				return []int{v}
			}
		}

	}

	return nil
}
