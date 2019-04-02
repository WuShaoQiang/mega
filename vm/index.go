package vm

import (
	"fmt"

	"github.com/WuShaoQiang/mega/model"
)

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username string) IndexViewModel {
	u1, err := model.GetUserByUsername(username)
	if err != nil {
		fmt.Println(err)
	}
	posts, err := model.GetPostsByUserID(u1.ID)
	if err != nil {
		fmt.Println(err)
	}
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts}
	v.SetCurrentUser(username)
	return v
}
