package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"dev-test/nubank-dev-test-2k21/app/dto/command"
)

func main() {
	var lines []interface{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())

		jsonData := []byte(scanner.Text())
		var line map[string]interface{}
		json.Unmarshal(jsonData, &line)

		if _, ok := line["account"]; ok {
			var accountLine command.AccountLine
			err := json.Unmarshal(jsonData, &accountLine)
			fmt.Println("aa")
			if err == nil {
				lines = append(lines, accountLine)
				continue
			}
		}

		if _, ok := line["transaction"]; ok {
			var transactionLine command.TransactionLine
			err := json.Unmarshal(jsonData, &transactionLine)
			if err == nil {
				lines = append(lines, transactionLine)
				continue
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("fail")
		os.Exit(-1)
	}

	fmt.Printf("%+v\n", lines)
}
