package model

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	LastSeen     *time.Time
	AboutMe      string `gorm:"type:varchar(140)"`
	Avatar       string `gorm:"type:varchar(200)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

// SetPassword func: Set PasswordHash
func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

// CheckPassword func return whether the password is true
func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

// GetUserByUsername func: Return User
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// SetAvatar func
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

// AddUser func
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	user.SetAvatar(email)
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return user.FollowSelf()
}

//UpdateUserByUsername can update user's content
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	return db.Model(item).Update(contents).Error
}

//UpdateLastSeen update the time of last_seen
func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

//UpdateAboutMe update user's aboutme
func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, contents)
}

// Follow someone
func (u *User) Follow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Append(u).Error
}

// Unfollow someone
func (u *User) Unfollow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Delete(u).Error
}

// FollowSelf follow myself
func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

//FollowersCount return the numbers of followers
func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}

//FollowingIDs return the userID who you are following
func (u *User) FollowingIDs() []int {
	var ids []int
	rows, err := db.Table("follower").Where("follower_id=?", u.ID).Select("user_id, follower_id").Rows()
	if err != nil {
		log.Println("Counting following failed : ", err)
		return ids
	}
	defer rows.Close()
	for rows.Next() {
		var id, followerID int
		rows.Scan(&id, &followerID)
		ids = append(ids, id)
	}
	return ids
}

// FollowingCount func return how many person you have followed for now
func (u *User) FollowingCount() int {
	ids := u.FollowingIDs()
	return len(ids)
}

//FollowingPosts return the posts of people who you follow
func (u *User) FollowingPosts() (*[]Post, error) {
	var posts []Post
	ids := u.FollowingIDs()
	err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

//IsFollowedByUser return whether the user has been followed by you
func (u *User) IsFollowedByUser(username string) bool {
	user, _ := GetUserByUsername(username)
	ids := user.FollowingIDs()
	for _, id := range ids {
		if u.ID == id {
			return true
		}
	}
	return false
}

//CreatePost can let you create a post with text(string)
func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}
