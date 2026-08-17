package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Create a reader to read from the terminal input
var reader = bufio.NewReader(os.Stdin)

// create temp state for exit
var exit bool = false

// Declare input variables
var purchasePrice float64
var downpaymentPercentage float64
var feePercentage float64
var monthlyInstallments int
var purchaseDate time.Time
var err error

// Declare calculations variables
var downpayment float64
var financeAmount float64
var fee float64
var totalRepayable float64
var installmentPrice float64

// Declare single installment
type Installment struct {
	Number  int
	DueDate time.Time
	Amount  float64
}

// Create a slice to act as your "temp place" for the schedule
var schedule []Installment

func clearConsole() {

	var cmd *exec.Cmd

	// This checks which operating system the program is running on
	if runtime.GOOS == "windows" {
		// If it's Windows, it uses the "cls" command
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// If it is NOT Windows (meaning it's Linux or Mac), it uses the "clear" command
		cmd = exec.Command("clear")
	}

	// Connect the command's output to the terminal
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getInputs() {

	// 1. Get Purchase Price
	fmt.Print("Enter purchase price: ")
	priceInput, _ := reader.ReadString('\n')
	priceInput = strings.TrimSpace(priceInput)
	purchasePrice, err = strconv.ParseFloat(priceInput, 64)
	if err != nil {
		fmt.Println("Invalid price format.")
		return
	}

	// 2. Get Fee Precentage
	fmt.Print("Enter Fee percentage (e.g., 20 for 20%):")
	feetpercentageInput, _ := reader.ReadString('\n')
	// Clean up any trailing newline or carriage return characters
	feetpercentageInput = strings.TrimSpace(feetpercentageInput)
	feePercentage, err = strconv.ParseFloat(feetpercentageInput, 64)
	if err != nil {
		fmt.Println("Invalid price format.")
		return
	}
	feePercentage = feePercentage / 100.0

	// 3. Get Downpayment Precentage
	fmt.Print("Enter downpayment percentage (e.g., 20 for 20%):")
	downpaymentpercentageInput, _ := reader.ReadString('\n')
	// Clean up any trailing newline or carriage return characters
	downpaymentpercentageInput = strings.TrimSpace(downpaymentpercentageInput)
	downpaymentPercentage, err = strconv.ParseFloat(downpaymentpercentageInput, 64)
	if err != nil {
		fmt.Println("Invalid price format.")
		return
	}
	downpaymentPercentage = downpaymentPercentage / 100.0

	// 4. Get Number of Installments
	fmt.Print("Enter number of monthly installments: ")
	monthsInput, _ := reader.ReadString('\n')
	monthsInput = strings.TrimSpace(monthsInput)
	monthlyInstallments, err = strconv.Atoi(monthsInput)
	if err != nil {
		fmt.Println("Invalid installment count format.")
		return
	}

	// 5. Get Purchase Date (format: YYYY-MM-DD)
	fmt.Print("Enter purchase date (YYYY-MM-DD): ")
	dateInput, _ := reader.ReadString('\n')
	dateInput = strings.TrimSpace(dateInput)
	purchaseDate, err = time.Parse("2006-01-02", dateInput)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}
}

func basicCalculations() {
	// Calculations
	downpayment = purchasePrice * downpaymentPercentage
	fee = purchasePrice * feePercentage
	financeAmount = purchasePrice - downpayment
	totalRepayable = financeAmount + fee
	installmentPrice = totalRepayable / float64(monthlyInstallments)

	//output
	fmt.Printf("Purchase Date: %s\n", purchaseDate.Format("2006-01-02"))
	fmt.Printf("Downpayment: %.2f\n", downpayment)
	fmt.Printf("Fee: %.2f\n", fee)
	fmt.Printf("Total Repayable: %.2f\n", totalRepayable)
	fmt.Printf("Installment Price: %.2f\n", installmentPrice)
}

func installmentsCalculation() {
	var accumulatedSum float64

	originalDay := purchaseDate.Day()
	startYear := purchaseDate.Year()
	startMonth := int(purchaseDate.Month())

	// --- PART 1: Generate and Store ---
	for i := 1; i <= monthlyInstallments; i++ {
		targetTotalMonths := startMonth + i - 1
		targetYear := startYear + (targetTotalMonths / 12)
		targetMonth := (targetTotalMonths % 12) + 1

		lastDayOfTargetMonth := getLastDayOfMonth(targetYear, time.Month(targetMonth))

		targetDay := originalDay
		if targetDay > lastDayOfTargetMonth {
			targetDay = lastDayOfTargetMonth
		}
		dueDate := time.Date(targetYear, time.Month(targetMonth), targetDay, 0, 0, 0, 0, purchaseDate.Location())

		var currentInstallment float64
		if i == monthlyInstallments {
			currentInstallment = totalRepayable - accumulatedSum
		} else {
			currentInstallment = math.Round(installmentPrice*100) / 100
			accumulatedSum += currentInstallment
		}

		// Store it in our temporary slice
		schedule = append(schedule, Installment{
			Number:  i,
			DueDate: dueDate,
			Amount:  currentInstallment,
		})

		// Print the standard schedule row
		fmt.Printf("%d: %s: Rs. %.2f\n", i, dueDate.Format("2006-01-02"), currentInstallment)
	}
}

func getLastDayOfMonth(year int, month time.Month) int {
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastDay := nextMonth.AddDate(0, 0, -1)
	return lastDay.Day()
}

func latepaymentCalculation() {
	for {
		fmt.Print("\nDo you want to calculate a late payment? (y/n) (x for exit): ")
		checkLate, _ := reader.ReadString('\n')
		checkLate = strings.TrimSpace(strings.ToLower(checkLate))

		if checkLate != "y" {
			if checkLate == "x" {
				fmt.Printf("Exiting ...")
				exit = true
				break
			}
			fmt.Printf("Returning ...")
			break
		}

		fmt.Print("Enter the installment number to check: ")
		instNumInput, _ := reader.ReadString('\n')
		instNum, err := strconv.Atoi(strings.TrimSpace(instNumInput))

		if err != nil || instNum < 1 || instNum > len(schedule) {
			fmt.Println("Invalid installment number.")
			continue
		}

		// Retrieve the specific installment from our stored slice
		targetInst := schedule[instNum-1]
		clearConsole()
		fmt.Printf("Installment %d is due on %s for Rs. %.2f.\n", targetInst.Number, targetInst.DueDate.Format("2006-01-02"), targetInst.Amount)
		fmt.Print("Enter actual payment date (YYYY-MM-DD): ")
		paidInput, _ := reader.ReadString('\n')
		paidDate, err := time.Parse("2006-01-02", strings.TrimSpace(paidInput))

		if err != nil {
			fmt.Println("Invalid date format.")
			continue
		}

		// Calculate Late Penalty using the stored data
		daysLate := int(paidDate.Sub(targetInst.DueDate).Hours() / 24)
		penaltyPercentage := 0.0

		if daysLate > 0 {
			if daysLate <= 7 {
				penaltyPercentage = 0.02
			} else {
				completeBlocks := (daysLate - 1) / 30
				penaltyPercentage = 0.05 + (float64(completeBlocks) * 0.01)
			}
		}

		penaltyAmt := targetInst.Amount * penaltyPercentage
		finalAmt := targetInst.Amount + penaltyAmt

		fmt.Printf("\n--- Late Payment Result ---\n")
		fmt.Printf("Due Date: %s | Paid Date: %s (%d days late)\n", targetInst.DueDate.Format("2006-01-02"), paidDate.Format("2006-01-02"), daysLate)
		fmt.Printf("Base Amount: Rs. %.2f\n", targetInst.Amount)
		fmt.Printf("Penalty Added: Rs. %.2f\n", penaltyAmt)
		fmt.Printf("Total Owed: Rs. %.2f\n", finalAmt)
		fmt.Println("---------------------------")
	}
}

func main() {
	exit = false

	fmt.Printf("Installment Scheduler \n\n")

	//Get inputs
	getInputs()

	for !exit {
		clearConsole()
		fmt.Printf("Installment Scheduler \n\n")

		// Schedule Info
		fmt.Printf("\n--- Schedule Info ---\n")
		basicCalculations()

		// Installments
		fmt.Printf("\n--- Installments ---\n")
		installmentsCalculation()

		// Late payment
		fmt.Printf("\n--- OPTIONAL: Late payment calculation ---\n")
		latepaymentCalculation()
	}

}
