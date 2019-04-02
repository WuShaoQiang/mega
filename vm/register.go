package vm

import (
	"log"

	"github.com/WuShaoQiang/mega/model"
)

// RegisterViewModel struct
type RegisterViewModel struct {
	LoginViewModel
}

// RegisterViewModelOp struct
type RegisterViewModelOp struct{}

// GetVM func
func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

// CheckUserExist func : true --> exist, false --> not exist
func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can't find user : ", username)
		return false
	}
	return true
}

// AddUser func
func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}
