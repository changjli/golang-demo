package atom_user

import (
	"net/http"
	"strconv"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func UpdateUserByIdV2(context *gin.Context) {
	var updateData atom_user.InputUserModel
	err := context.ShouldBindJSON(&updateData)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(context.Param("id"))

	getData, err := atom_user.UpdateUserByIdUseCaseV2(updateData, id)

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
			"message": "Success Insert Data User",
		})
	}
}
