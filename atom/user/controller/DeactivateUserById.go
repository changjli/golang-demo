package atom_user

import (
	"net/http"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func DeactivateUserById(context *gin.Context) {
	var updateData atom_user.UpdateUserModel
	context.ShouldBindJSON(&updateData)

	getData, err := atom_user.DeactiveUserByUseCase(updateData)

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
			"message": "Success Update Data User",
		})
	}
}
