package vm

import "github.com/WuShaoQiang/mega/model"

// ProfileViewModel struct
type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
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
	v.SetCurrentUser(sUser)

	return v, nil
}
