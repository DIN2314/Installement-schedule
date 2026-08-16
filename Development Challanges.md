# Development Challenges

## 🎯 Hardest Aspects of Development

### 1. **Month-End Date Handling** ⭐⭐⭐ (Highest Difficulty)
The most complex challenge was correctly calculating due dates when the purchase day falls on a date that doesn't exist in some months (e.g., the 31st in February or April).

**Problem:** If someone purchases on January 31st with a monthly schedule, what's the due date for February? There is no February 31st.

**Solution:** Implemented `getLastDayOfMonth()` to check the maximum day in the target month and intelligently clamp the date:
```go
lastDayOfTargetMonth := getLastDayOfMonth(targetYear, time.Month(targetMonth))
targetDay := originalDay
if targetDay > lastDayOfTargetMonth {
    targetDay = lastDayOfTargetMonth
}
```

### 2. **Floating-Point Precision in Calculations** ⭐⭐⭐
Financial calculations require exact precision to the rupee.

**Problem:** Dividing the total amount by number of installments can result in long decimal values. Distributing rounding errors across installments is challenging.

**Solution:** Used rounding on each installment and corrected the final payment:
```go
if i == monthlyInstallments {
    currentInstallment = totalRepayable - accumulatedSum  // Last payment covers remainder
} else {
    currentInstallment = math.Round(installmentPrice*100) / 100  // Round to 2 decimals
}
```

### 3. **Cross-Platform Console Clearing** ⭐⭐
Windows uses `cls` while Unix-based systems use `clear`. Had to implement OS detection.

**Solution:** Used `runtime.GOOS` to detect the operating system and execute the appropriate command.

### 4. **Late Payment Penalty Tier Logic** ⭐⭐
Implementing the tiered penalty system required careful mathematical calculation:
- 1-7 days: 2%
- 8-30 days: 5%
- 31+ days: 5% + 1% per 30-day block

**Challenge:** Correctly calculating "complete blocks" for penalties beyond 30 days without off-by-one errors.

### 5. **Input Validation and Error Recovery** ⭐⭐
Ensuring the program handles invalid inputs gracefully without crashing.

**Challenge:** Need to validate dates, numbers, and percentages while maintaining a good user experience.

### 6. **State Management** ⭐
Managing global variables for the installment schedule and ensuring data persists across function calls without side effects.

## 🔮 Future Improvements

- Refactor global variables into a struct-based configuration
- Add support for custom penalty tiers
- Implement export to CSV/PDF
- Add unit tests for edge cases
- Create a web interface version