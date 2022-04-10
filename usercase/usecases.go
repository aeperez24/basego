package usercase

import (
	"aeperez24/basego/dto"
	"aeperez24/basego/port"
)

type UserCaseError struct {
	error
	Code    int
	Message string
}

//USER
type UserUsercase struct {
	UserService port.UserService
}

func (userUserCase UserUsercase) CreateUser(user dto.UserWithPasswordDto) UserCaseError {
	basicUser := userUserCase.UserService.FindBasicUser(user.Username)
	if basicUser != (dto.BasicUserDto{}) {
		return UserCaseError{nil, 400, "username already exists"}
	}
	userByDocument := userUserCase.UserService.FindBasicUserByDocument(user.Username)
	if userByDocument != (dto.BasicUserDto{}) {
		return UserCaseError{nil, 400, "DocumentId already exists already exists"}

	}
	err := userUserCase.UserService.CreateUser(user)

	if err != nil {
		return UserCaseError{err, 500, "error creating user"}

	}
	return UserCaseError{}
}

//ANOTHER USECASE
