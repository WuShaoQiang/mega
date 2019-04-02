package vm

import (
	"fmt"

	"github.com/WuShaoQiang/mega/model"
)

// ProfileViewModel struct
type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	Editable    bool
	ProfileUser model.User
}

type ProfileViewModelOp struct{}

func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, err := model.GetPostsByUserID(u1.ID)
	if err != nil {
		return v, err
	}
	v.ProfileUser = *u1
	v.Posts = *posts
	v.Editable = (sUser == pUser)
	fmt.Println(v.Editable)
	v.SetCurrentUser(sUser)

	return v, nil
}
