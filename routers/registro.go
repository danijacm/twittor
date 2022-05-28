package routers

import (
	"encoding/json"
	"net/http"

	"github.com/danijacm/twittor/bd"
	"github.com/danijacm/twittor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t) // El body es un Stream, lo que quiere decir que despues de leerlo la primera vez, desaparece de la memoria

	if err != nil {
		http.Error(w, "Error en la recepcion de datos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Error, el email es requerido", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Error, el password debe ser de almenos 6 caracteres", 400)
	}

	_, encontrado, _ := bd.YaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
	}

	_, status, err := bd.InserttoRegistro(t)

	if err != nil {
		http.Error(w, "No fue posible insertar la info de usuario"+err.Error(), 500)
	}

	if status == false {
		http.Error(w, "No se logro hacer el insert en la BD"+err.Error(), 500)
	}

	w.WriteHeader(http.StatusCreated)
}
