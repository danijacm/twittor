package middlew

import (
	"net/http"

	"github.com/danijacm/twittor/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "No hay conexión con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
