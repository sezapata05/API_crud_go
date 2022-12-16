package commands

import (
	"bufio"
	"fmt"
	"go_cli/expenses"
	"log"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	fmt.Print("-> ")

	str, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\n", "", 1)

	return str, nil
}

func ShowInConsole(expenses_list []float32) {
	fmt.Println(contentString(expenses_list))
}

func Export(filename string, lis []float32) error {
	log.Println(filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(lis))
	if err != nil {
		return err
	}

	return w.Flush()
}

func contentString(expenses_list []float32) string {

	builder := strings.Builder{}

	max, min, avg, total := expensesDetails(expenses_list)

	fmt.Println("")

	for i, expense := range expenses_list {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))

		if i == len(expenses_list)-1 {
			fmt.Println("================================")
			fmt.Println("================================")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
			fmt.Println("================================")
		}
	}
	return builder.String()
}

func expensesDetails(expenses_list []float32) (max, min, avg, total float32) {

	if len(expenses_list) == 0 {
		return
	}

	min = expenses.Min(expenses_list...)
	max = expenses.Max(expenses_list...)
	avg = expenses.Average(expenses_list...)
	total = expenses.Sum(expenses_list...)

	return

}
