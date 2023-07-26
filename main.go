package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Expense struct {
	name  string
	value int
}

func main() {
	expenses := []Expense{}
	shouldStop := false
	reader := bufio.NewReader(os.Stdin)
	expensesTotal := 0

	for !shouldStop {
		fmt.Println("--- New expense --- \n")

		fmt.Println("Expense Name: ")
		expenseName, expenseNameErr := reader.ReadString('\n')
		logError(expenseNameErr)

		fmt.Println("Expense Value: ")
		expenseValueStr, expenseValueErr := reader.ReadString('\n')
		logError(expenseValueErr)

		expenseValue, err := strconv.Atoi(strings.TrimSpace(expenseValueStr))
		logError(err)

		expenses = append(expenses, Expense{
			value: expenseValue,
			name:  expenseName,
		})

		fmt.Println("Stop the program (y/n)")
		newShouldStop, newShouldStopErr := reader.ReadString('\n')
		
		logError(newShouldStopErr)
		if newShouldStop = strings.TrimSpace(newShouldStop); newShouldStop == "y" {
			shouldStop = true
		} else {
			shouldStop = false
		}
	}

	for i := 0; i < len(expenses); i++ {
		fmt.Println("Expense named %s has the value of %d", expenses[i].name, expenses[i].value)
		expensesTotal += expenses[i].value
	}

	fmt.Println("You had a total of: ", expensesTotal)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
