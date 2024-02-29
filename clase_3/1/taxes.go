package taxes

func Taxes(salary int) (tax int) {
	switch {
	case salary > 150000:
		tax += percentage(salary, 10)
		fallthrough
	case salary > 50000:
		tax += percentage(salary, 17)
	}
	return
}

func percentage(value int, percentage int) int {
	return value * percentage / 100
}
