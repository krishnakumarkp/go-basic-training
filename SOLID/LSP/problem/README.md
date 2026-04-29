💥 Why This Violates LSP
❌ Broken Assumption

Account says:

“All accounts can deposit AND withdraw”

But reality:

Fixed deposits cannot withdraw

❌ Substitution Fails
ProcessAccount(sa) // ✅ OK
ProcessAccount(fd) // 💥 crashes

👉 You cannot safely substitute FixedDepositAccount

🚨 Real Problems Caused
💥 Runtime crashes (panic)
🤥 Fake implementations (“not supported”)
🔗 Tight, incorrect abstraction
🧪 Hard to test safely
🧠 Misleading API design
🧠 Root Cause

You modeled behavior as universal when it is not