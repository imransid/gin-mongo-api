package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"fmt"
)

func CheckUserType(c *gin.Context, role string)(err error){
	userType := c.GetString("user_type")

	err = nil

	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}


// for check same user id
func MatchUserTypeToUid(c *gin.Context, userId string)(err error) {
	userType := c.GetString("userId")

	uid := c.GetString("uid")


	fmt.Printf("uid", userId )

	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}


