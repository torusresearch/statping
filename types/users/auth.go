package users

import (
	"fmt"
	"time"

	"github.com/statping/statping/utils"
)

// AuthUser will return the User and a boolean if authentication was correct.
// accepts username, and password as a string
func AuthUser(username, passwordHash string) (*User, bool) {
	user, err := FindByUsername(username)
	if err != nil {
		log.Warnln(fmt.Errorf("user %v not found", username))
		return nil, false
	}
	if utils.CheckHash(passwordHash, user.Password) {
		user.UpdatedAt = time.Now().UTC()
		user.Update()
		return user, true
	}
	return nil, false
}
