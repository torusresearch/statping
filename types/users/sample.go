package users

import (
	"github.com/torusresearch/statping/types/null"
)

func Samples() error {
	log.Infoln("Inserting Sample Users...")
	u2 := &User{
		Username: "testadmin",
		Password: "password123",
		Email:    "info@betatude.com",
		Admin:    null.NewNullBool(true),
	}

	if err := u2.Create(); err != nil {
		return err
	}

	u3 := &User{
		Username: "testadmin2",
		Password: "password123",
		Email:    "info@adminhere.com",
		Admin:    null.NewNullBool(true),
	}

	if err := u3.Create(); err != nil {
		return err
	}

	return nil
}
