package eight

import "io/ioutil"

// I need a file to hold the huge test data so that my IDE stops freezing
var testInput = ""

func loadTestData() {
	dat, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	testInput = string(dat)
}
