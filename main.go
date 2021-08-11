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

func getUsers( c echo.Context ) error {

	users , err := user_service.Read()
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	return c.JSON(http.StatusOK,users)
}
func createUser(c echo.Context) error {
	data := echo.Map{}
	err := c.Bind(&data);

	if  err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	oid := primitive.NewObjectID()

	user := models.User{
		ID: oid,
		Name: fmt.Sprintf("%v", data["nombre"]),
		Email: fmt.Sprintf("%v", data["correo"]),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var errService error = user_service.Create(user)

	if errService != nil {
		return c.JSON(http.StatusInternalServerError,errService.Error())
	}
	return c.JSON(http.StatusOK,user)
}
func updateUser(c echo.Context) error {
	data := echo.Map{}
	err := c.Bind(&data)

	if  err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	user := models.User{
		Name: fmt.Sprintf("%v", data["nombre"]),
		Email: fmt.Sprintf("%v", data["correo"]),
		UpdatedAt: time.Now(),
	}
	var errService error = user_service.Update(user,fmt.Sprintf("%v", data["id_user"]))
	if errService != nil{
		return c.JSON(http.StatusInternalServerError,errService.Error())
	}

	return c.JSON(http.StatusOK,user)
}
func deleteUser(c echo.Context) error {
	data := echo.Map{}
	err := c.Bind(&data)

	if  err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	var errService error = user_service.Delete(fmt.Sprintf("%v", data["id_user"]))
	if errService != nil{
		return c.JSON(http.StatusInternalServerError,errService.Error())
	}

	return c.JSON(http.StatusOK,map[string]interface{}{"mensaje":"Usuario eliminado"})
}




func main() {
	e := echo.New()

	e.GET("/api/user",getUsers)
	e.POST("/api/user",createUser)
	e.PUT("/api/user",updateUser)
	e.DELETE("/api/user",deleteUser)


	e.Start(":8080")
}