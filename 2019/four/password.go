package four

import "math"

// Checker takes a password and returns a true if it is valid. 
type Checker func(p int) bool

// ValidPasswordInRange returns the number of valid passwords in the given range. 
func ValidPasswordInRange(min, max int) int {
	count := 0
	for i := min; i <= max; i++ {
		if checkIsValid(i, checkIsSixDigits, checkRepeatingDigits, checkDoesNotDecrease ) {
			count++
		}
	}
	return count
}

func checkIsValid(password int, checkers ...Checker) bool {
	for _, c := range checkers {
		if !c(password) {
			return false
		}
	}
	
	return true
}

func checkDoesNotDecrease(p int) bool {
	// Going from left to right, the digits never decrease; 
	// they only ever increase or stay the same (like 111123 or 135679).
	last := -1
	for i := 6; i > 0; i-- {
		a := digit(p, i)
		if a < last {
			return false
		}
		last = a
	}
	return true
}

func checkIsSixDigits(p int) bool {
	if p < 100000 || p >= 1000000 {
		return false
	}
	return true
}

// assumes number coming in is 6digits long. Also this is ugly
func checkRepeatingDigits(number int) bool {
	valid := false
	for i := 6; i > 1; i-- {
		a := digit(number, i)
		b := digit(number, (i - 1))

		if a == b {
			c := -1
			if i > 2 {
				c = digit(number, (i - 2))
			}

			if a == c {
				// Jump to where i is
				i-- // -> b
				i-- // -> c
				i-- // after c
				for i > 1 {
					c = digit(number, i)
					if c == a {
						i-- // still the same batch
					} else {
						i++ // Undo the last skip
						break
					}
				}
			} else {
				valid = true
				break
			}

		}
	}

	return valid
}

// https://stackoverflow.com/questions/46753736/extract-digits-at-a-certain-position-in-a-number
func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
