# 🟢 What is a `rune` in Go?

* A **`rune` is an alias for `int32`**
* It represents a **Unicode code point** (a single character)
* Go uses **UTF-8 encoding**, so characters can take **1–4 bytes**

👉 In short:

> `rune = one Unicode character`

---

# 🔤 Example 1: Single Rune

```go
var value rune
value = '♬'
fmt.Printf("Character: %c , Decimal: %d,  Unicode: %U", value, value, value)
```

### 🔍 What happens here:

* `'♬'` is a **rune literal**
* Internally stored as a Unicode number

Output will look like:

```
Character: ♬ , Decimal: 9836, Unicode: U+266C
```

👉 Same value, different representations:

* `%c` → character
* `%d` → decimal value
* `%U` → Unicode format

---

# 🌏 Example 2: String with Unicode (`नमस्ते`)

```go
str := "नमस्ते"
```

This is a **UTF-8 string**, not a rune.

---

## 🔁 Iterating over runes

```go
for _, r := range str {
	fmt.Printf("Character: %c , Decimal: %d,  Unicode: %U", r, r, r)
}
```

### 🔍 Important:

* `range` over a string **automatically decodes UTF-8 into runes**
* Each `r` is a **rune (Unicode character)**

👉 Output (simplified):

```
न म स ् त े
```

Each printed separately with its Unicode value.

---

## 📏 Length: Bytes vs Runes

```go
len(str)
```

👉 Returns **number of bytes**, NOT characters

Example:

```
Length of the string (in bytes): 18
```

---

```go
len([]rune(str))
```

👉 Converts string → slice of runes
👉 Returns **number of characters**

Example:

```
Length of the string (in runes): 6
```

---

# ⚠️ Key Insight

| Concept            | Meaning                      |
| ------------------ | ---------------------------- |
| `string`           | Sequence of **bytes**        |
| `rune`             | Single **Unicode character** |
| `len(str)`         | Byte length                  |
| `len([]rune(str))` | Character count              |

---

# 🧠 Why `rune` matters

Without runes:

* Unicode strings (Hindi, emoji, etc.) break
* Indexing a string gives **bytes**, not characters

Example:

```go
str[0] // ❌ gives byte, not full character
```

Correct way:

```go
[]rune(str)[0] // ✅ gives full character
```

---

# ✅ Summary

* `rune` = Unicode character (`int32`)
* Strings = UTF-8 encoded bytes
* Use `range` or `[]rune()` to safely handle characters
* Essential for **internationalization (i18n)** and **text processing**