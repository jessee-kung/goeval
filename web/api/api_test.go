package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserGetResponse struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type APITestSuite struct {
	suite.Suite

	router *gin.Engine
}

func (suite *APITestSuite) SetupTest() {
	suite.router = setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/register/api_ut01", nil)
	suite.router.ServeHTTP(w, req)
}

func (suite *APITestSuite) TearDownTest() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/unregister/api_ut01", nil)
	suite.router.ServeHTTP(w, req)

	w = httptest.NewRecorder() // Remove for TestUserRegisterRouter
	req, _ = http.NewRequest("GET", "/user/unregister/api_ut02", nil)
	suite.router.ServeHTTP(w, req)
}

func (suite *APITestSuite) TestUserGetRouter() {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/get/api_ut01", nil)
	suite.NoError(err, "Unexpected error occurred %v", err)

	suite.router.ServeHTTP(w, req)
	var body []byte
	body, err = ioutil.ReadAll(w.Result().Body)
	suite.NoError(err, "Unexpected error occurred %v", err)

	var response UserGetResponse
	err = json.Unmarshal(body, &response)
	suite.NoError(err, "Unexpected error occurred %v", err)

	suite.Equal(200, w.Code)
	suite.Equal("api_ut01", response.Name)
}

func (suite *APITestSuite) TestUserUnregisterRouter() {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/unregister/api_ut01", nil)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.router.ServeHTTP(w, req)
	var body []byte
	body, err = ioutil.ReadAll(w.Result().Body)
	if err != nil {
		suite.T().Fatal(err)
	}

	var response UserGetResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		suite.T().Fatal(err)
	}
	
	suite.Equal(200, w.Code)
	suite.Equal("api_ut01", response.Name)

	req, err = http.NewRequest("GET", "/user/get/api_ut01", nil)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.router.ServeHTTP(w, req)

	body, err = ioutil.ReadAll(w.Result().Body)
	if err != nil {
		suite.T().Fatal(err)
	}
	
	suite.Equal(404, w.Code)
}

func (suite *APITestSuite) TestUserRegisterRouter() {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/register/api_ut02", nil)
	suite.NoError(err, "Unexpected error occurred %v", err)

	suite.router.ServeHTTP(w, req)
	w.Result()

	req, err = http.NewRequest("GET", "/user/get/api_ut02", nil)
	suite.NoError(err, "Unexpected error occurred %v", err)

	suite.router.ServeHTTP(w, req)

	var body []byte
	body, err = ioutil.ReadAll(w.Result().Body)
	suite.NoError(err, "Unexpected error occurred %v", err)

	var response UserGetResponse
	err = json.Unmarshal(body, &response)
	suite.NoError(err, "Unexpected error occurred %v", err)

	suite.Equal(200, w.Code)
	suite.Equal("api_ut02", response.Name)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
