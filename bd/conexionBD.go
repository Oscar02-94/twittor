package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conxion ala BD*/
var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://sandra:canelito@cluster0.b2thy.mongodb.net/appgo?retryWrites=true&w=majority")

/*ConectarBD es la funcion quwe nos permite conectar la bd*/
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
	log.Println("Conexion Exitosa con la base de datos")
	return client
}

/*ChequeoConnection es el pin ala base de datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
