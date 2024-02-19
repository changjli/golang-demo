package atom_user

import (
	"net/http"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func GetAllData(context *gin.Context) {
	getData, err := atom_user.GetAllDataUserUseCase()

	if err != nil {
		// context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		// context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"data":    getData,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    getData,
			"message": "Success Get Data User",
		})
	}
}
