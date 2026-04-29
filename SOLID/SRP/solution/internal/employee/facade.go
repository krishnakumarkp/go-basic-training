package employee

type EmployeeFacade struct {
	payCalc  PayCalculator
	reporter HourReporter
	saver    EmployeeSaver
}

func NewEmployeeFacade(
	p PayCalculator,
	r HourReporter,
	s EmployeeSaver,
) *EmployeeFacade {
	return &EmployeeFacade{
		payCalc:  p,
		reporter: r,
		saver:    s,
	}
}

func (f EmployeeFacade) CalculatePay(emp Employee) float64 {
	return f.payCalc.Calculate(emp)
}

func (f EmployeeFacade) ReportHours(emp Employee) {
	f.reporter.Report(emp)
}

func (f EmployeeFacade) SaveEmployee(emp Employee) {
	f.saver.Save(emp)
}
