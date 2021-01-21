package bd

import (
	"context"
	"time"

	"github.com/Oscar02-94/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExixteUsuario recibe un email de parametro y checa si ya esta en la base de datos*/
func ChequeoYaExixteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuario")

	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
