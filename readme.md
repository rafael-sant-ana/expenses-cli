# Expense-Tracker CLI

A simple CLI in go to practice my concepts of file management and go syntax.

Initial scope: 
1. Storing the data in a json `expenses.json`

## Commands

**add**
```sh
expense add --amount 32.50 --category food --note "lunch"
```
**list**
```sh
expense list
```

**summary**
```sh
expense summary
```

## How to run
Example:
```sh
go run . add --amount 25.5 --category food --note "hello"
```
