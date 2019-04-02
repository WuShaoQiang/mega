package vm

import (
	"github.com/WuShaoQiang/mega/model"
)

// ProfileViewModel struct
type ProfileViewModel struct {
	BaseViewModel
	Posts          []model.Post
	Editable       bool
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	ProfileUser    model.User
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
	//Not him/herself
	if !v.Editable {
		v.IsFollow = u1.IsFollowedByUser(sUser)
	}
	v.FollowersCount = u1.FollowersCount()
	v.FollowingCount = u1.FollowingCount()
	v.SetCurrentUser(sUser)

	return v, nil
}

func Follow(sUser, pUser string) error {
	u, err := model.GetUserByUsername(sUser)
	if err != nil {
		return err
	}
	return u.Follow(pUser)
}

func UnFollow(sUser, pUser string) error {
	u, err := model.GetUserByUsername(sUser)
	if err != nil {
		return err
	}
	return u.Unfollow(pUser)
}
