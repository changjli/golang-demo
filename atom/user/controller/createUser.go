package atom_user

import (
	"net/http"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var inputData atom_user.InputUserModel
	// validasi data type
	err := context.ShouldBindJSON(&inputData)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	getData, err := atom_user.InsertUserUseCase(inputData)

	// validasi query
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
