package bd

import "golang.org/x/crypto/bcrypt"

func EcriptarPassword(pass string) (string, error) {
	costo := 8 // Entre mas grande sea el costo, mas seguro será el cifrado del password y mas demora el algoritmo en devolver el valor cifrado
	passEnBytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(passEnBytes), err
}
