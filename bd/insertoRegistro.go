package bd

import (
	"context"
	"time"

	"github.com/danijacm/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InserttoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // al final de la función candela el contexto que tiene el timeout

	db := MongoCx.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EcriptarPassword(u.Password) // Se cifra el password que viene desde el cliente antes de guardarlo en la BD

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID) // Obtengo el ID del registro insertado
	return ObjID.String(), true, nil
}
