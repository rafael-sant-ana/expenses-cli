package main

import (
	"fmt"
	"os"
	"flag"
	"expenses-tracker/expense"
	"expenses-tracker/store"
	"github.com/google/uuid"
	"time"
)

var expensesStore store.Store

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Execute it with `add`, `list` or `summary`")
		return
	}
	var err error

	expensesStore, err = store.NewStore("./store.json")

	if err != nil {
		fmt.Println("erro carregando")
		return
	}

	switch args[0] {
		case "add":
			addExpenses(args[1:])
		case "list":
			listExpenses(args[1:])
		case "summary":
			summaryExpenses(args[1:])
	}
}

func addExpenses(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	amount :=  addCmd.Float64("amount", 0, "Amount transactioned")
	category := addCmd.String("category", "", "Category of the transaction")
	notes := addCmd.String("note", "", "Additional information of expense")

	addCmd.Parse(args)

	newExpense := expense.Expense{
		uuid.New(),
		*amount,
		*category,
		*notes,
		time.Now(),
	}

	expensesStore.Add(newExpense)
}

func listExpenses(args []string) {
	fmt.Println(expensesStore.List())
}

func summaryExpenses(args []string) {

}
