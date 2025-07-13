package controller

import (
	"fmt"

	"github.com/Kayuan165/CRUD-Golang/src/configuration/rest_err"
	"github.com/Kayuan165/CRUD-Golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		restErr := rest_err.NewBadRequestErr(
			fmt.Sprintf("There are some invalid fields in the request, error %s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
