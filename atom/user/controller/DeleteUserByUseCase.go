package atom_user

import (
	"net/http"
	"strconv"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func DeleteUserById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	getData, err := atom_user.DeleteUserByUseCase(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"data":    getData,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"data":    getData,
			"message": "Success Insert Data User",
		})
	}
}
