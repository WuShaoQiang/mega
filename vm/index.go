package vm

import (
	"fmt"

	"github.com/WuShaoQiang/mega/model"
)

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
	u1, err := model.GetUserByUsername("rene")
	if err != nil {
		fmt.Println(err)
	}
	posts, err := model.GetPostsByUserID(u1.ID)
	if err != nil {
		fmt.Println(err)
	}
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	return v
}
