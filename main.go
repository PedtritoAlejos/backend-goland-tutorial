package main

import (
	"Usuario/models"
	user_service "Usuario/services/user.service"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func createUser(c echo.Context) error {
	data := echo.Map{}

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
	}
	oid := primitive.NewObjectID()

	user := models.User{
		ID: oid,
		Name: fmt.Sprintf("%v", data["nombre"]),
		Email: fmt.Sprintf("%v", data["correo"]),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user_service.Create(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,user)
}

func getUsers( c echo.Context ) error {

	users , err := user_service.Read()
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	return c.JSON(http.StatusOK,users)
}

func main() {
	e := echo.New()

	e.GET("/api/user",getUsers)
	e.POST("/api/user",createUser)


	e.Start(":8080")
}