package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCx = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@myfirstcluster.rjwds.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD Funcion para conectar con la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD...")
	return client
}

func CheckConnection() int {
	err := MongoCx.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
