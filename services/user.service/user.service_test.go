package user_service_test

import (
	"Usuario/models"
	user_service "Usuario/services/user.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var userId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()


	user := models.User{
		ID: oid,
		Name: "Juan",
		Email: "juan@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err:= user_service.Create(user)

	if err!= nil {
		t.Error("La prueba de persistencia de datos a fallado")
		t.Fail()
	}else{
		 t.Log("La prueba finalizo con exito!")
	}
}
func TestRead(t *testing.T) {

	users , err := user_service.Read()

	if err != nil {
		t.Error("Se ha presentado un error en la consulta de usuarios")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("La consulta no retorno datos")
		t.Fail()
	}else {
		t.Log("La prueba finalizo con exito ")
	}
}

func TestUpdate(t *testing.T) {
	user := models.User{
		Name: "Goku",
		Email: "goku@gmail.com",

	}

	err := user_service.Update(user,userId)

	if err != nil {
		t.Error("Error al tratar de actualizar el usuario")
		t.Fail()
	}else {
		t.Log("La prueba de actualización finalizó con exito")
	}
}

/*
func TestDelete(t *testing.T) {

	err := user_service.Delete(userId)

	if err != nil {
		t.Error("Error al eliminar el usuario ")
		t.Fail()
	}else {
		t.Log("La prueba de eliminación finalizó con exito")
	}


}

 */

