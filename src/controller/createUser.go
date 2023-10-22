package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lechitz/GoMongoConnect/src/configuration/rest_err"
	"github.com/lechitz/GoMongoConnect/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error = %s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}