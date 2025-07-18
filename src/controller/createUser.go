package controller

import (
	"fmt"

	"github.com/Kayuan165/CRUD-Golang/src/configuration/validation"
	"github.com/Kayuan165/CRUD-Golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
