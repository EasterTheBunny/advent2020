package data

// ComputeExpenseReport finds two numbers in the data source that add up to
// 2020 and returns the product of the two numbers.
func ComputeExpenseReport(d []int) int {
	product := 0

	for i, e := range d {
		for j, f := range d {
			if j == i {
				break
			}

			if e+f == 2020 {
				product = e * f
			}
		}
	}

	return product
}
