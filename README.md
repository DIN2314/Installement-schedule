# 📅 Installment Scheduler

A powerful command-line tool for calculating installment payment schedules with late payment penalties. Perfect for managing purchases financed through monthly installments.

## ✨ Features

- **Flexible Payment Calculations** - Input purchase price, downpayment percentage, and fees
- **Automatic Schedule Generation** - Generate monthly installment schedules with accurate due dates
- **Smart Date Handling** - Intelligently handles month-end dates (e.g., Feb 29, 31st day months)
- **Late Payment Penalties** - Calculate penalties for late payments with tiered penalty structure
- **Multi-Platform Support** - Works seamlessly on Windows, Linux, and macOS
- **Interactive CLI** - User-friendly terminal interface with real-time calculations

## 🚀 Quick Start

### Running the Application

Pre-built binaries are available in the `./bin/` directory for multiple platforms. No installation required!

**On Linux/macOS:**
```bash
./bin/installment-schedule-linux-amd64    # For Linux (AMD64)
./bin/installment-schedule-darwin-amd64   # For macOS (Intel)
./bin/installment-schedule-darwin-arm64   # For macOS (Apple Silicon)
```

**On Windows:**
Simply download the appropriate binary and run it directly.

### Building from Source

**On Windows:**
```powershell
./build.ps1
```

**On Linux/macOS:**
If you have PowerShell Core (7+) installed:
```bash
pwsh ./build.ps1
```

Otherwise, use Go directly:
```bash
go build -o installment-schedule cmd/main.go
./installment-schedule
```

## 📖 Usage

1. **Enter Purchase Price** - The total cost of the item
2. **Enter Fee Percentage** - Processing/financing fee (e.g., 20 for 20%)
3. **Enter Downpayment Percentage** - Initial payment percentage (e.g., 10 for 10%)
4. **Enter Number of Installments** - How many monthly payments
5. **Enter Purchase Date** - Start date in YYYY-MM-DD format

The program will display:
- Downpayment amount
- Total fee charged
- Total repayable amount
- Individual installment payment schedule

### Late Payment Calculation

After generating the schedule, you can optionally calculate late payment penalties:
- Payments up to 7 days late: **2% penalty**
- Payments 8-30 days late: **5% penalty**
- Payments beyond 30 days: **5% + 1% per 30-day block**

## 📊 Example

```
Enter purchase price: 100000
Enter Fee percentage: 20
Enter downpayment percentage: 10
Enter number of monthly installments: 12
Enter purchase date: 2026-08-17

--- Schedule Info ---
Purchase Date: 2026-08-17
Downpayment: 10000.00
Fee: 18000.00
Total Repayable: 162000.00
Installment Price: 13500.00

--- Installments ---
1: 2026-09-17: Rs. 13500.00
2: 2026-10-17: Rs. 13500.00
...
```

## 🛠️ Architecture

- **getInputs()** - Collects user input with validation
- **basicCalculations()** - Computes financial metrics
- **installmentsCalculation()** - Generates the payment schedule
- **latepaymentCalculation()** - Handles penalty calculations
- **clearConsole()** - Cross-platform terminal clearing

## 📝 License

This project is licensed under the MIT License.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.