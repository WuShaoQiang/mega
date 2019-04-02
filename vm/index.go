package vm

import (
	"fmt"

	"github.com/WuShaoQiang/mega/model"
)

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func return a specific IndexViewModel
func (IndexViewModelOp) GetVM(username, flash string) IndexViewModel {
	u1, err := model.GetUserByUsername(username)
	if err != nil {
		fmt.Println(err)
	}
	posts, err := model.GetPostsByUserID(u1.ID)
	if err != nil {
		fmt.Println(err)
	}
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts, flash}
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
