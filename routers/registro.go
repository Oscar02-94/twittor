package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Oscar02-94/twittor/bd"
	"github.com/Oscar02-94/twittor/models"
)

/*Registro es la funcion para crear en la base de datos el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos resividos"+err.Error(), 400)
		return

	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar un contraseña de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExixteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya exixte un usuario registrado con ese Email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado Insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
