package employee

import "fmt"

// Employee struct holds data (this part is fine)
type Employee struct {
	Name  string
	Hours float64
	Rate  float64
}

// BUSINESS LOGIC (Accounts Department)
// If salary rules change (tax, overtime, bonuses),
// the Accounts team will require changes here
func (e Employee) CalculatePay() float64 {
	return e.Hours * e.Rate
}

// REPORTING (HR Department)
// If HR wants a different report format (PDF, dashboard, logs),
// this method must change
func (e Employee) ReportHours() {
	fmt.Printf("%s worked %.2f hours\n", e.Name, e.Hours)
}

// PERSISTENCE (Database Admin / Backend Team)
// If database changes (SQL → NoSQL, API, ORM, etc.),
// this method must change
func (e Employee) SaveEmployee() {
	fmt.Printf("saving %s to database \n", e.Name)
}
