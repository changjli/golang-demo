package atom_user

import (
	"net/http"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func GetUserByUsernameV2(context *gin.Context) {
	username := context.Param("username")

	getData, err := atom_user.GetUserByUsernameCaseV2(username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"data":    getData,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    getData,
			"message": "Success Get User Data by Username",
		})
	}
}
