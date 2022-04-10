package port

import (
	"aeperez24/basego/dto"
	"net/http"
)

type UserService interface {
	CreateUser(user dto.UserWithPasswordDto) error
	ValidateUserameAndPassword(username string, password string) bool
	FindBasicUser(username string) dto.BasicUserDto
	FindBasicUserByDocument(document string) dto.BasicUserDto
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
