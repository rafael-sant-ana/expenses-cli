package store

import (
	"encoding/json"
	"expenses-tracker/expense"
	"fmt"
	"os"
)

type Store struct {
	Expenses []expense.Expense
	filePath string
}

func NewStore(filePath string) (Store, error) {
	store := Store{
		Expenses: []expense.Expense{},
		filePath: filePath,
	}

	store.loadFromFile()

	return store, nil	
}

func (s Store) Add(expense expense.Expense) {
	s.Expenses = append(s.Expenses, expense)
	
	s.writeInFile()
}

func (s Store) List() []expense.Expense {
	return s.Expenses
}

func (s Store) writeInFile() { // (s Store) faz nao ter mutacao pq n tem ponteiro
	data, err := json.Marshal(s.Expenses)
	if err != nil{
		fmt.Println("erro json parseando")
		return	
	}

	os.WriteFile(s.filePath, data, 0644) // 0644 -> Permissao
}

func (s *Store) loadFromFile() { // usa * pra mutar
	data, err := os.ReadFile(s.filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("error while opening the file")
		return
	}

	err = json.Unmarshal(data, &s.Expenses)
}
