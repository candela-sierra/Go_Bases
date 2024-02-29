package salary

type category string

const (
	A category = "A"
	B category = "B"
	C category = "C"
)

const (
	hourlySalaryC = 1000
	hourlySalaryB = 1500
	hourlySalaryA = 3000
)

func SalaryByCategory(category category, minutes int) (salary int) {
	hours := minutesToHours(minutes)
	switch category {
	case C:
		salary = salaryCalculator(hours, hourlySalaryC, 0)
		// hours * hourlySalaryA
	case B:
		salary = salaryCalculator(hours, hourlySalaryB, 20)
		//  hours * hourlySalaryB * 120/100
	case A:
		salary = salaryCalculator(hours, hourlySalaryA, 50)
		// hours * hourlySalaryC * 150/100
	}
	return
}

func salaryCalculator(hours int, hourlySalary int, monthlyBonus int) int {
	return hours * hourlySalary * (100 + monthlyBonus) / 100
}

func minutesToHours(minutes int) int {
	return minutes / 60
}
