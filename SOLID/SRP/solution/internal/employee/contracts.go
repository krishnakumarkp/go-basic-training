package employee

type PayCalculator interface {
	Calculate(Employee) float64
}

type HourReporter interface {
	Report(Employee)
}

type EmployeeSaver interface {
	Save(Employee)
}
