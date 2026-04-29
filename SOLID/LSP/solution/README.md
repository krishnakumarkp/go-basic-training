Why This Design Is Correct
✔ No fake implementations
✔ No panic / runtime surprises
✔ Compiler enforces correct usage
✔ Each type implements only what it supports
🧠 Final Insight

Model capabilities, not assumptions

Instead of:

“All accounts can withdraw” ❌

We say:

“Some accounts can withdraw” ✅
“All accounts can deposit” ✅
🔥 One-Line Takeaway

If a behavior is not universal, don’t put it in the base abstraction