package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
	"github.com/jessee-kung/goeval/common/model/user"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("user/get/:name", func(c *gin.Context) {
		u := user.GetUser(c.Param("name"))
		if u != nil {
			c.JSON(http.StatusOK, gin.H {
				"status": http.StatusOK,
				"uuid": u.UUID,
				"name": u.Name,
			})
		} else {
			c.JSON(http.StatusOK, gin.H {
				"status": http.StatusNotFound,
			})
		}
	})

	r.GET("user/register/:name", func(c *gin.Context) {
		u := user.RegisterUser(c.Param("name"))
		if u != nil {
			c.JSON(http.StatusOK, gin.H {
				"status": http.StatusOK,
				"uuid": u.UUID,
				"name": u.Name,
			})
		} else {
			c.JSON(http.StatusOK, gin.H {
				"status": http.StatusBadRequest,
			})
		}
	})

	r.GET("user/unregister/:name", func(c *gin.Context) {
		u := user.UnregisterUser(c.Param("name"))
		if u != nil {
			c.JSON(http.StatusOK, gin.H {
				"status": http.StatusOK,
				"uuid": u.UUID,
				"name": u.Name,
			})
		} else {
			c.JSON(http.StatusOK, gin.H {
				"status": http.StatusNotFound,
			})
		}
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}