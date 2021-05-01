package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/service"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func Run() {
	operations := input.NewOperations()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		jsonData := []byte(scanner.Text())
		var line map[string]interface{}
		json.Unmarshal(jsonData, &line)

		if _, ok := line[input.LINE_ACCOUNT]; ok {
			var accountLine input.AccountLine
			err := json.Unmarshal(jsonData, &accountLine)
			if err == nil {
				operations.AddLine(accountLine)
			}
			continue
		}

		if _, ok := line[input.LINE_TRANSACTION]; ok {
			var transactionLine input.TransactionLine
			err := json.Unmarshal(jsonData, &transactionLine)
			if err == nil {
				operations.AddLine(transactionLine)
			}
			continue
		}
	}

	if scanner.Err() != nil {
		fmt.Println("fail on read stdin")
		os.Exit(-1)
	}

	validators := []validator.ValidatorInterface{
		validator.NewCardActiveValidator(),
		validator.NewCardLimitValidator(),
		validator.NewHighTransactionsValidator(2, 120),
		validator.NewDoubleTransactionValidator(120),
	}

	authorizeService := service.NewAuthorizeService(validators)
	authorizeService.HandleOperations(operations)
}
