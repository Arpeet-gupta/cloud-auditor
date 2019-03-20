package environment

import (
	"os/user"
)

func GetUserHomeDir() (string, error) {
	myUser, userError := user.Current()
	if userError != nil {
		return "", userError
	}
	path := myUser.HomeDir

	return path, nil
}
