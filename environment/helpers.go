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

func UniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
}
