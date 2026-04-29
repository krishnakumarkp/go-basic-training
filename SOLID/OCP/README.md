---

# 📘 OCP Solution Using Interfaces (Go)

## ✅ Principle (from SOLID principles)

> Open for extension, closed for modification

---

## 🧠 Core Idea

* Define an **interface** (`PaymentMethod`)
* `PaymentProcessor` depends on the **interface**, not concrete types
* New behavior = **new implementation**, not modifying old code

---

## 🏗️ Design

### 1. Interface (Extension Point)

```go id="n1"
type PaymentMethod interface {
	Pay(amount float64)
}
```

---

### 2. Stable Core (Closed for Modification)

```go id="n2"
type PaymentProcessor struct{}

func (p PaymentProcessor) Process(pm PaymentMethod, amount float64) {
	pm.Pay(amount)
}
```

👉 This code **never changes**

---

### 3. Extend via New Implementations

```go id="n3"
type UPI struct{}

func (u UPI) Pay(amount float64) {
	fmt.Println("Processing UPI payment")
}
```

👉 Add new types, don’t modify existing code

---

## 🔁 How Extension Works

* Add new struct implementing `PaymentMethod`
* Pass it to `Process()`
* ✅ No changes to `PaymentProcessor`

---

## ⚖️ Design Choice (DI vs Parameter)

### Option A: Pass as parameter (Flexible)

```go id="n4"
Process(pm PaymentMethod, amount float64)
```

* One processor → many payment types
* Runtime flexibility

---

### Option B: Inject dependency (Fixed)

```go id="n5"
type PaymentProcessor struct {
	method PaymentMethod
}
```

* One processor → one payment type
* Better for configured services

---

## 🧠 Key Takeaway

> Interface = extension point
> New feature = new implementation
> ❌ Never modify stable core code

---
