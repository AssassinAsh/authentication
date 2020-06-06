package components

import (
	"authentication/repo/slave"
	"log"
)

//CheckUsernameExistence -
func CheckUsernameExistence(username string) bool {
	slaveRepo := slave.NewUserCredentialsSlaveRepo()

	_, err := slaveRepo.ReadUserCredentialsByUsername(username)

	if err != nil {
		log.Println(err)
		return true
	}

	return false
}
