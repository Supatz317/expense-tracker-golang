package model

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Expense struct {
	ID          int `json:"id"` // ID is auto incremented
	Description string
	Amount      float64
	Date        time.Time
}

type Expenses []Expense

// map month number to month name
var monthNames = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func (Exps *Expenses) List() {

	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("# ID", "Date", "Description", "Amount")

	for index, t := range *Exps {
		table.AddRow(strconv.Itoa(index), t.Date.Format("2006-01-02"), t.Description, strconv.Itoa(int(t.Amount)))
	}
	table.Render()
}

func (Exps *Expenses) AddExpense(description string, amount float64) {
	expense := Expense{
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
	// fmt.Printf("Expense added successfully : \n%s => %.2f baht", description, amount)
	*Exps = append(*Exps, expense)
	_ = Exps

}


func (Exps *Expenses) DeleteExpense(index int) error{
	e := *Exps
	if err := e.validateIndex(index); err != nil {
		return err 
	}

	*Exps = append(e[:index], e[index+1:]...)

	fmt.Println("Expense deleted successfully")
	return nil
}

func (Exps *Expenses) validateIndex(index int) error {
	if index < 0 || index >= len(*Exps) {
		err := errors.New("invalid index")
		return err
	}

	return nil
}

func (Exps *Expenses) UpdateExpense(index int, description string, amount float64) error {
	e := *Exps
	if err := e.validateIndex(index); err != nil {
		return err 
	}

	e[index].Description = description
	e[index].Amount = amount

	fmt.Println("Expense updated successfully")
	return nil
}

func (Exps *Expenses) Summary() {
	sum := 0.0
	for _, t := range *Exps {
		sum += t.Amount
	}
	fmt.Printf("Total expenses: %.2f baht\n", sum)
}

func (Exps *Expenses) SummaryByMonth(month int) {
	sum := 0.0

	for _, t := range *Exps {
		// fmt.Println("=========", t.Date , month)
		
		if int(t.Date.Month()) == month {
			sum += t.Amount
		}
	}
	m := monthNames[month]
	fmt.Printf("Total expenses for %s: %.2f baht\n", m ,sum)
}

var Exps = Expenses{}