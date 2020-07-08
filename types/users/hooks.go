package users

import "github.com/torusresearch/statping/utils"

func (u *User) BeforeCreate() error {
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA256Hash()
	u.ApiSecret = utils.NewSHA256Hash()
	return nil
}
