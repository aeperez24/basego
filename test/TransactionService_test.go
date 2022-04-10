package test

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
	"fmt"
	"testing"
)

func TestGetTransactionsWithSuccess(t *testing.T) {
	repo := AccountRepositoryMock{}
	repo.FindAccountByAccountNumberFn = func(account string) model.Account {
		transactionList := make([]model.Transaction, 0)
		transactionList = append(transactionList, model.Transaction{})
		transactionList = append(transactionList, model.Transaction{})
		if account == "account" {
			return model.Account{AccountNumber: account, Transactions: transactionList}
		} else {
			return model.Account{}
		}
	}
	accountNumber := "account"
	service := services.NewTransactionService(repo)
	transactions, _ := service.GetTransactions(accountNumber)
	if len(transactions) != 2 {
		t.Errorf("expected %v and received %v", 2, len(transactions))
	}

}

func TestSaveTransferTransaction(t *testing.T) {
	repo := AccountRepositoryMock{}
	transactionMap := make(map[string]model.Transaction)

	repo.SaveTransactionFn = func(account string, transaction model.Transaction) error {
		transactionMap[account] = transaction
		return nil
	}
	service := services.NewTransactionService(repo)
	transactionDto := dto.TransactionDto{AccountFrom: "from", AccountTo: "to", Amount: 100, Type: "Transfer"}
	transaction := model.Transaction{AccountFrom: "from", AccountTo: "to", Amount: 100, Type: model.TransferType}

	service.SaveTransaction(transactionDto)

	if transaction != transactionMap["from"] {
		t.Errorf("expected %v and received %v", transaction, transactionMap["from"])
	}

	if transaction != transactionMap["to"] {
		t.Errorf("expected %v and received %v", transaction, transactionMap["to"])
	}

}

func TestSaveDepositTransaction(t *testing.T) {
	repo := AccountRepositoryMock{}
	transactionMap := make(map[string]model.Transaction)

	repo.SaveTransactionFn = func(account string, transaction model.Transaction) error {
		transactionMap[account] = transaction
		return nil
	}
	service := services.NewTransactionService(repo)
	transactionDto := dto.TransactionDto{AccountFrom: "from", AccountTo: "to", Amount: 100, Type: "Deposit"}
	transaction := model.Transaction{AccountTo: "to", Amount: 100, Type: model.DepositType}

	service.SaveTransaction(transactionDto)

	if transaction != transactionMap["to"] {
		t.Errorf("expected %v and received %v", transaction, transactionMap["to"])
	}

}

func TestSaveWithdrawTransaction(t *testing.T) {
	repo := AccountRepositoryMock{}
	transactionMap := make(map[string]model.Transaction)

	repo.SaveTransactionFn = func(account string, transaction model.Transaction) error {
		transactionMap[account] = transaction
		return nil
	}
	service := services.NewTransactionService(repo)
	transactionDto := dto.TransactionDto{AccountFrom: "from", AccountTo: "to", Amount: 100, Type: "Withdraw"}
	transaction := model.Transaction{AccountFrom: "from", Amount: 100, Type: model.WithdrawType}

	service.SaveTransaction(transactionDto)

	if transaction != transactionMap["from"] {
		t.Errorf("expected %v and received %v", transaction, transactionMap["to"])
	}

}
func TestSaveTransactionError(t *testing.T) {
	repo := AccountRepositoryMock{}
	transactionMap := make(map[string]model.Transaction)

	repo.SaveTransactionFn = func(account string, transaction model.Transaction) error {
		transactionMap[account] = transaction
		return nil
	}
	service := services.NewTransactionService(repo)
	transactionDto := dto.TransactionDto{AccountFrom: "from", AccountTo: "to", Amount: 100, Type: "BadType"}

	saveError := service.SaveTransaction(transactionDto)
	fmt.Println(saveError)

	if saveError == nil {
		t.Error("expected Error")
	}

}
