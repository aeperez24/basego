package port

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"net/http"
)

type AccountService interface {
	GetBalance() (float32, error)
	TransferMoneyTo(accountNumber string, amount float32) error
	Deposit(amount float32) error
	Withdraw(amount float32) error
}

type UserService interface {
	CreateUser(user dto.UserWithPasswordDto) error
	ValidateUserameAndPassword(username string, password string) bool
	FindBasicUser(username string) dto.BasicUserDto
	FindBasicUserByDocument(document string) dto.BasicUserDto
}

type TransactionService interface {
	GetTransactions(string) ([]model.Transaction, error)
	SaveTransaction(dto.TransactionDto) error
}

type AccountHandler interface {
	GetBalance(http.ResponseWriter, *http.Request)
	TransferMoney(http.ResponseWriter, *http.Request)
	Deposit(http.ResponseWriter, *http.Request)
	GetTransactions(http.ResponseWriter, *http.Request)
}
type UserHandler interface {
	CreateUser(http.ResponseWriter, *http.Request)
}
type AuthenticationHandler interface {
	Authenticate(http.ResponseWriter, *http.Request)
}
type Server interface {
	Start() error
	Stop()
}
type TokenService interface {
	CreateToken(dto.BasicUserDto) (string, error)
	ExtractBasicUseDtoFromToken(receivedToken string) (dto.BasicUserDto, error)
}

const (
	TransferType string = "Transfer"
	DepositType  string = "Deposit"
	WithdrawType string = "Withdraw"
)

type context_key string

const (
	LoggedUserKey context_key = "loggedUser"
)
