package one

// CalculateFuelForModule returns the needed fuel for a module
func CalculateFuelForModule(mass int) int {
	p := mass / 3

	f := p - 2

	if f <= 0 {
		return 0
	}
	return f + CalculateFuelForModule(f)
}

// CalculateFuelForModules returns the needed fuel for all modules
func CalculateFuelForModules(modules []int) int {
	totalFuel := 0
	for _, m := range modules {
		totalFuel += CalculateFuelForModule(m)
	}
	return totalFuel
}
