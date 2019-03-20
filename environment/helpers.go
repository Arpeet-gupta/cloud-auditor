package environment

import (
	"os/user"
)

func GetUserHomeDir() (string, error) {
	myUser, userError := user.Current()
	
}