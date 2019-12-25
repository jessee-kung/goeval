package user

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
    suite.Suite
}

func (suite *UserTestSuite) SetupTest() {
	RegisterUser("ut01")
	RegisterUser("ut02")
	RegisterUser("ut03")
}

func (suite *UserTestSuite) TearDownTest() {
	UnregisterUser("ut01")
	UnregisterUser("ut02")
	UnregisterUser("ut03")
}

func (suite *UserTestSuite) TestRegisterDuplicateUser() {
	var testUsername string = "ut01"
	u := RegisterUser(testUsername)
	suite.Nil(u, "user name '%s' is expected to be registered failed", testUsername)
}

func (suite *UserTestSuite) TestRegisterNonExistUser() {
	var testUsername string = "ut04"
	uReg := RegisterUser(testUsername)
	suite.NotNil(uReg, "user name '%s' is expected to be registered success", testUsername)
	suite.Equal(uReg.Name, testUsername, "user name is expected to be '%s'", testUsername)

	uGet := GetUser(testUsername)
	suite.NotNil(uGet, "user name '%s' is expected to be registered success", testUsername)
	suite.Equal(uGet.Name, testUsername, "user name is expected to be '%s'", testUsername)
}

func (suite *UserTestSuite) TestUnregisterExistUser() {
	testUsernames := []string{"ut01", "ut02", "ut03"}
	for _, testUsername := range(testUsernames) {
		u := UnregisterUser(testUsername)
		suite.NotNil(u, "user name '%s' is expected to be removed succcess", testUsername)
	}
}

func (suite *UserTestSuite) TestUnregisterNonExistUser() {
	var testUsername string = "ut05"
	uReg := UnregisterUser(testUsername)
	suite.Nil(uReg, "user name '%s' is expected to be removed failed due to not exist", testUsername)
}

func TestUserTestSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}