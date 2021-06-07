package routes

import (
	"go-enm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)


func UserRegister(c *gin.Context){
	user := models.User{}

	err := c.ShouldBindJSON(&user)

	if(err !=nil){
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.Register(&conn)

	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := user.GetAuthToken()
	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"email": token,
	})
}