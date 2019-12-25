package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jessee-kung/goeval/common/model/logger"
)

// User is the structure which keeps the user information
type User struct {
	UUID string
	Name string
}

var (
	mUsers = make(map[string]*User)
)


// RegisterUser add a new user into the registered user list.
// The user instance will be returned if success, nil otherwise
func RegisterUser (name string) *User{
	_, ok := mUsers[name]
	if ok {
		fmt.Println(name)
		logger.Warning("'%s' is already exist", name)
		return nil
	}
	id, err := uuid.NewUUID()
	if err != nil {
		logger.Error("User register for '%s' failed, UUID generation failed", name)
		return nil
	}

	mUsers[name] = &User{Name: name, UUID: id.String()}
	return mUsers[name]
}

// UnregisterUser remove the specific user from the registered user list.
// The user instance will be returned if success, nil otherwise
func UnregisterUser (name string) *User {
	user, ok := mUsers[name]
	if ok != true {
		logger.Warning("'%s' is not exist", name)
		return nil
	}
	delete(mUsers, name)
	return user
}

// GetUser retrieve the specific user from the registered user list.
// The user instance will be returned if success, nil otherwise
func GetUser (name string) *User{
	user, ok := mUsers[name]
	if ok != true {
		logger.Warning("'%s' is not exist", name)
		return nil
	}
	return user
}