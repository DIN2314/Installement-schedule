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

### 📦 Getting Started
Pre-built binaries are available in the `./bin/` directory for immediate use:
```bash
./bin/installment-schedule-linux-amd64    # For Linux (AMD64)
./bin/installment-schedule-darwin-amd64   # For macOS (Intel)
./bin/installment-schedule-darwin-arm64   # For macOS (Apple Silicon)
```

No installation required! Or, build from source:

**On Windows:**
```powershell
./build.ps1
```

**On Linux/macOS:**
If you have PowerShell Core (7+):
```bash
pwsh ./build.ps1
```

Otherwise, use Go directly:
```bash
go build -o installment-schedule cmd/main.go
./installment-schedule
```

### 👨‍💻 Contributors
- Dinesh Weerasinghe

---

**Release Date:** August 17, 2026  
**Status:** Stable