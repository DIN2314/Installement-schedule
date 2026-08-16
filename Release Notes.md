# Release Notes

## Version 1.0.0 - Initial Release (2026-08-17)

### 🎉 Features
- ✅ Complete installment schedule generation
- ✅ Accurate payment calculations with downpayment and fees
- ✅ Intelligent date handling for month-end scenarios
- ✅ Interactive late payment penalty calculator
- ✅ Cross-platform support (Windows, macOS, Linux)
- ✅ User-friendly command-line interface

### 🔧 Technical Details
- Built with Go for fast, reliable performance
- Supports currency formatting with two decimal places
- Real-time schedule calculations
- Robust input validation with error messages

### 📋 How to Use
1. Run the application
2. Enter financial details (price, percentages, installments, date)
3. Review the generated payment schedule
4. Optionally calculate late payment penalties

### ⚠️ Known Limitations
- Currencies are displayed in PKR (Rupees) - customization not yet supported
- Penalty structure is fixed - custom tiers not supported
- No data persistence between sessions
- Limited to monthly installment periods

### 🐛 Bug Fixes
- None (Initial Release)

### 📦 Installation
```bash
go build -o installer cmd/main.go
./installer
```

### 📝 What's Next
- v1.1.0: Add CSV export functionality
- v1.2.0: Support for custom penalty tiers
- v1.3.0: Database persistence
- v2.0.0: Web-based interface

### 👨‍💻 Contributors
- Development Team

---

**Release Date:** August 17, 2026  
**Status:** Stable