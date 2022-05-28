package bd

import (
	"context"
	"time"

	"github.com/danijacm/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func YaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCx.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email} // Formateo la condición que voy a usar para la busqueda al formato de la BD

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex() // Saco el ID del registro encontrado y lo paso a string
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
