
SRP isn’t just about “number of functions,” it’s about **who forces the code to change**. In your case, *different stakeholders* (different departments) can all demand changes to the same struct. That’s a textbook violation of the idea behind SOLID principles.

---

## 🧾 Your Code with Stakeholder-Based SRP Explanation

```go id="k2h8mz"
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
```

---

## 🎯 The Core SRP Violation (Now Crystal Clear)

Your `Employee` struct has **multiple stakeholders**:

* 💰 **Accounts team** → cares about `CalculatePay`
* 🧑‍💼 **HR team** → cares about `ReportHours`
* 🗄️ **DB/Admin team** → cares about `SaveEmployee`

👉 Each team can independently request changes.

---

## 💥 Why That’s Dangerous

* One struct becomes a **battlefield of changes**
* Teams unintentionally **break each other’s logic**
* You get:

  * Merge conflicts
  * Regression bugs
  * Slower development

---

## 🧠 SRP Reframed (Best Way to Remember)

> A module should have **one reason to change = one stakeholder**

Your code has:

> ❌ One module (`Employee`) → **three stakeholders**
